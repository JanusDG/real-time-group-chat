package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JanusDG/real-time-group-chat/odt"
	"github.com/gorilla/websocket"
)

type Server struct {
	Port     int
	Debug_on bool
	Upgrader websocket.Upgrader
	Counter  int
	Connections []UserConnection
}

type UserConnection struct {
	Id 		int
	Conn 	*websocket.Conn
}

func NewUserConnection(id int,conn *websocket.Conn) *UserConnection{
	return &UserConnection{Id: id, Conn: conn}
}

// func Init - initializer for server instance
func (s *Server) Init(port int, DEBUG_ON bool) {
	s.Port = port
	s.Debug_on = DEBUG_ON
}

// func NewServer - constructor for server instance
func NewServer(port int, debug_on bool) *Server {
	return &Server{Port: port, Debug_on: debug_on, Counter: 1,
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}, Connections: make([]UserConnection, 0)}
}

func (s *Server) reader(conn *websocket.Conn) {
	for {
		var m = comms.Message{}
		err := conn.ReadJSON(&m)
		if err != nil {
			log.Println("Error reading json.", err)
		}
		log.Printf("Got message: %#v\n", m)

		err = s.Connections[1].Conn.WriteJSON(m)
		if err != nil {
			log.Println(err)
		}
		// read in a message
		// messageType, p, err := conn.ReadMessage()
		// if err != nil {
		// 	log.Printf("func: reader, error in message read, %s", err)
		// 	return
		// }

		// log.Println("user said " + string(p))

		// // write it back
		// if err := conn.WriteMessage(messageType, []byte("u sent me \""+string(p)+"\"")); err != nil {
		// 	log.Printf("func: reader, error in message write, %s", err)
		// 	return
		// }
		

	}
}

func (s *Server) wsEndpoint(w http.ResponseWriter, r *http.Request) {

	s.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client #%s Connected", s.Counter)

	// Init User
	var init = comms.NewInitUser(s.Counter)
	log.Println(init.Id)
	// err = ws.WriteJSON(init)
	// if err != nil {
	// 	log.Println(err)
	// }
	var newUser = NewUserConnection(s.Counter, ws)
	s.Connections = append(s.Connections, *newUser)
		
		
	// TODO make mutex here
	s.Counter = s.Counter + 1

	// err = ws.WriteMessage(1, []byte("Hi Client!"))
	// if err != nil {
	// 	log.Println(err)
	// }

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	s.reader(ws)
}

func (s *Server) RunServer() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/websocket.html")
	})
	http.HandleFunc("/ws", s.wsEndpoint)

	http.HandleFunc(
		"/hello",
		func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "hello") })

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.Port), nil))
}
