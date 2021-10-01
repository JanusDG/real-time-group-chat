package client

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
	"bufio"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/JanusDG/real-time-group-chat/odt"
	// "github.com/satori/go.uuid"
)

type Client struct {
	Id string
}

func NewClient() *Client{
	return &Client{Id: ""}
}

func (c *Client) Init() {
	var addr = flag.String("addr", "localhost:8080", "http service address")

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			if (c.Id == ""){
				var init = comms.InitUser{}
				err := conn.ReadJSON(&init)
				if err != nil {
					log.Println("Error reading json.", err)
				}else{
					c.Id = init.Id
				}
			}

			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Printf("func: reader, error in message read, %s", err)
				return
			}else{
				log.Printf(string(p))
			}

			var m = comms.Message{}
			err = conn.ReadJSON(&m)
			if err != nil {
				log.Println("Error reading json.", err)
			}else{
				// c.Id = init.Id
				// log.Printf("Got message: %#v\n", m)
				log.Printf("Got message: %s\n", m.Content)
			}

			
		}
	}()

	// inputInit := make(chan []string)
	inputInit := make(chan string)
	defer close(inputInit)
	
	finished := make(chan bool)
	defer close(finished)
	go func() {
		// for {			
		log.Println("What's ur name?")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		// var split = strings.Split(text, " ")
		// inputInit<-split
		inputInit<-text
		finished <- true
		return
		// }
	}()
	
	input := make(chan string)
	defer close(input)
	go func (finished chan bool){
		<- finished
		for {
			log.Println("Write message")
			text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			// inputInit<-split
			input<-text
		}
		return 
	}(finished)
	

	for {
		select {
		case <-done:
			return
		case in := <-inputInit:
			in = strings.TrimSpace(in)
			err := conn.WriteMessage(websocket.TextMessage, []byte(in))
			if err != nil {
				log.Println("write:", err)
				return
			}

		case inn := <-input:
			// TODO handle the case if id is not defined yet
			// if user connected 
			var split = strings.Split(inn, " ")
			if len(split) != 2 {
				log.Println("Invalid input")
				return 
			}
			
			log.Printf("u typed %s, to %s", strings.TrimSpace(split[1]), strings.TrimSpace(split[0]))
			// if (false){
			var message = comms.NewMessage(c.Id, strings.TrimSpace(split[0]), strings.TrimSpace(split[1]))
			
			var err = conn.WriteJSON(message)
			if err != nil {
				log.Println(err)
			}
			// }
			

		case <-interrupt:
			log.Println("interrupt")

			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

