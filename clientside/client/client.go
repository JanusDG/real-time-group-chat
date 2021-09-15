package client

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
	"bufio"

	"github.com/gorilla/websocket"
	"github.com/JanusDG/real-time-group-chat/odt"

)

type Client struct {
	Id int
}

func NewClient() *Client{
	return &Client{Id: -1}
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
			var m = comms.Message{}
			err := conn.ReadJSON(&m)
			if err != nil {
				log.Println("Error reading json.", err)
			}
			log.Printf("Got message: %#v\n", m)

			// var m = comms.InitUser{}
			// err := conn.ReadJSON(&m)
			// if err != nil {
			// 	log.Println("Error reading json.", err)
			// }
			// log.Printf("Got message: %#v\n", m)
			// c.Id = m.Id
		}
	}()

	input := make(chan string)
	defer close(input)

	go func() {
		for {
			text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			input<-text
		}
	}()


	for {
		select {
		case <-done:
			return
		case in := <-input:
			// TODO handle the case if id is not defined yet
			// TODO no hardcode
			var message = comms.NewMessage(1, 2, in)

			var err = conn.WriteJSON(message)
			if err != nil {
				log.Println(err)
			}
			// err := conn.WriteMessage(websocket.TextMessage, []byte(in))
			// if err != nil {
			// 	log.Println("write:", err)
			// 	return
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

