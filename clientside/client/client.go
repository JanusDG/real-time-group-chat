package client

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
	"bufio"
	// "json"

	"github.com/gorilla/websocket"
	"github.com/JanusDG/real-time-group-chat/transfer"

)

type Client struct {
	
}

func NewClient() *Client{
	return &Client{}
}

func (*Client) Init() {
	var addr = flag.String("addr", "localhost:8080", "http service address")

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			var m = comms.Mess{}
			err := c.ReadJSON(&m)
			// if err != nil {
			// 	log.Println("read:", err)
			// 	return
			// }
			if err != nil {
				log.Println("Error reading json.", err)
			}
			log.Printf("Got message: %#v\n", m)
			// log.Printf("recv: %s", message)
			// var t = json.Unmarshal()
			// log.Printf(message)
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


	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case in := <-input:
			err := c.WriteMessage(websocket.TextMessage, []byte(in))
			if err != nil {
				log.Println("write:", err)
				return
			}

		// case t := <-ticker.C:
		// 	err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
		// 	if err != nil {
		// 		log.Println("write:", err)
		// 		return
		// 	}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
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

