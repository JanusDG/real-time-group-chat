package client

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/JanusDG/real-time-group-chat/odt"
	"github.com/JanusDG/real-time-group-chat/clientside/client/front"
)

type Client struct {
	Id string
	Webserver front.WebServer
	WgAuth *sync.WaitGroup
	WgContacts *sync.WaitGroup
	WgSendMessage *sync.WaitGroup
	WgReadMessage *sync.WaitGroup
}

func NewClient() *Client{
	var wgA = &sync.WaitGroup{}
	wgA.Add(1)
	var wgC = &sync.WaitGroup{}
	wgC.Add(1)
	var wgS = &sync.WaitGroup{}
	wgS.Add(1)
	var wgR = &sync.WaitGroup{}
	wgR.Add(1)
	return &Client{Id: "", WgAuth: wgA,WgReadMessage:wgR, WgContacts:wgC, WgSendMessage:wgS, Webserver: *front.NewWebServer(8082, wgA, wgC, wgS, wgR)}
}

func (c *Client) Init() {
	go c.Webserver.RunWeb()

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
				var init = comms.InitId{}
				err := conn.ReadJSON(&init)
				if err != nil {
					log.Println("Error reading json.", err)
				}else{
					c.Id = init.Id
				}
			}

			var uo = comms.UsersOption{}
			err = conn.ReadJSON(&uo)
			c.Webserver.Contacts = &uo.Users
			c.WgContacts.Done()

			var m = comms.Message{}
			err = conn.ReadJSON(&m)
			if err != nil {
				log.Println("Error reading json.", err)
			}else{
				var recieveMap = make(map[string]string)
				recieveMap["from"] = m.From
				recieveMap["content"] = m.Content
				c.Webserver.RecieveData = &recieveMap
				c.WgReadMessage.Done()
			}

			
		}
	}()

	inputInit := make(chan string)
	defer close(inputInit)
	
	finished := make(chan bool)
	defer close(finished)
	go func() {
		inputInit<-"text"
		finished <- true
		return
	}()
	
	input := make(chan string)
	defer close(input)
	go func (finished chan bool){
		<- finished
		c.WgContacts.Wait()
		input<-"text"
		// }
		return 
	}(finished)
	

	for {
		select {
		case <-done:
			return
		case <-inputInit:
			c.WgAuth.Wait()
			var m = c.Webserver.GetUserData()
			for k, v := range m {
				var message = comms.NewInitUser(k, v)
				
				var err = conn.WriteJSON(message)
				if err != nil {
					log.Println(err)
				}
			}

		case <-input:
			
			c.WgSendMessage.Wait()

			var messageMap = c.Webserver.MessageData

			var message *comms.Message
			for k, v := range messageMap {
				message = comms.NewMessage(c.Id, k, v)

			}
			
			var err = conn.WriteJSON(message)
			if err != nil {
				log.Println(err)
			}
			

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
