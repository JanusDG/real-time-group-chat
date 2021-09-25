package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JanusDG/real-time-group-chat/odt"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"

)

type Server struct {
	DB *comms.DataBase
	Port     int
	Debug_on bool
	Upgrader websocket.Upgrader
	ConnectionsMap map[uuid.UUID]*websocket.Conn
	UserMap map[uuid.UUID]comms.User
	// Connections []UserConnection
}

// type UserConnection struct {
// 	Id 		int
// 	Conn 	*websocket.Conn
// }

// func NewUserConnection(id int,conn *websocket.Conn) *UserConnection{
// 	return &UserConnection{Id: id, Conn: conn}
// }

// func Init - initializer for server instance
func (s *Server) Init(port int, DEBUG_ON bool) {
	s.Port = port
	s.Debug_on = DEBUG_ON
}

// func NewServer - constructor for server instance
func NewServer(database *comms.DataBase, port int, debug_on bool) *Server {
	return &Server{DB: database, Port: port, Debug_on: debug_on,
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}, 
		ConnectionsMap: make(map[uuid.UUID]*websocket.Conn),
		UserMap: make(map[uuid.UUID]comms.User),
		}
}

func (s *Server) readerLogin(conn *websocket.Conn,uniqueId uuid.UUID) {
	for (!s.UserMap[uniqueId].Loginned) {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("func: reader, error in message read, %s", err)
			return
		}

		var u = s.UserMap[uniqueId]
		u.Name = string(p)
		u.Loginned = true
		s.UserMap[uniqueId] = u
		
		log.Println("New user with name: " + string(p))

		// write it back
		// if err := conn.WriteMessage(messageType, []byte("u sent me \""+string(p)+"\"")); err != nil {
		// 	log.Printf("func: reader, error in message write, %s", err)
		// 	return
		// }
		

	}
	return
}

func (s *Server) writerNewInitUser(conn *websocket.Conn) uuid.UUID {
	var new_uuid = uuid.NewV1()
	var init = comms.NewInitUser(new_uuid)
	log.Println("New Client Connected")
	log.Println(init.Id)
	var err = conn.WriteJSON(init)
	if err != nil {
		log.Println(err)
	}
	var newUser = comms.NewUser("")
	newUser.Id = new_uuid
	s.UserMap[new_uuid] = *newUser
	return new_uuid
}

// TODO make writing of map with name:key
func (s *Server) writerContacts(conn *websocket.Conn, name string) {
	// var ConstactsMap = make(map[uuid.UUID]) 
	var message = ""
	for _, group := range s.DB.Groups{
		for _, member := range group.Members{
			if member.Name == name {
				message += "Available group to write: " + group.Name + "\n"
			}
		}
	}
	for _, user := range s.DB.Users{
		if user.Name != name {
			message += "Available user to write: " + user.Name + "\n"
		}
	}
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write:", err)
		return
	}
	return
}

// TODO add broadcast to group
func (s *Server) redirectMesasage(conn *websocket.Conn) {
	for {
		var m = comms.Message{}
		err := conn.ReadJSON(&m)
		if err != nil {
			log.Println("Error reading json.", err)
		}
		log.Printf("Got message: %#v\n", m)

		// all this mess will dissapear once server will be connected to bd
		for _, group := range s.DB.Groups {
			if group.Name == m.To {
				for _, member := range group.Members {
					for key, user := range s.UserMap{
						if user.Name == member.Name {
							err = s.ConnectionsMap[key].WriteJSON(m)
							log.Println("Sended back")
							
							if err != nil {
								log.Println(err)
							}
						}
					}
					// log.Println(user.Name)
					// log.Println(user.Id)
				}
				return
			}
		}
		for key, element := range s.UserMap {
			if m.To == element.Name {
				// log.Printf(element.Name)
				// log.Printf(s.ConnectionsMap[key])
				err = s.ConnectionsMap[key].WriteJSON(m)
				log.Println("Sended back")
				if err != nil {
					log.Println(err)
				}
				return
			}
		}
	}
}


func (s *Server) wsEndpoint(w http.ResponseWriter, r *http.Request) {

	s.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// s.Counter++
	
	// Init User
	var new_uuid = s.writerNewInitUser(ws)
	s.ConnectionsMap[new_uuid] = ws
	s.readerLogin(ws, new_uuid)
	
	s.writerContacts(ws, s.UserMap[new_uuid].Name)

	s.redirectMesasage(ws)
	// var newUser = NewUserConnection(s.Counter, ws)
	// s.Connections = append(s.Connections, *newUser)
		

	// err = ws.WriteMessage(1, []byte("Hi Client!"))
	// if err != nil {
	// 	log.Println(err)
	// }

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
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
