package server

import (
	// "fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JanusDG/real-time-group-chat/serverside/server/database"
	"github.com/JanusDG/real-time-group-chat/odt"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"

)

type Server struct {
	DB *database.Database
	Port     int
	Debug_on bool
	Upgrader websocket.Upgrader
	UserMap map[string]comms.User
}


// func Init - initializer for server instance
func (s *Server) Init(port int, DEBUG_ON bool) {
	s.Port = port
	s.Debug_on = DEBUG_ON
}

// func NewServer - constructor for server instance
func NewServer(database *database.Database, port int, debug_on bool) *Server {
	return &Server{DB: database, Port: port, Debug_on: debug_on,
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}, 
		UserMap: make(map[string]comms.User),
		}
}

func (s *Server) writerNewInitId(conn *websocket.Conn) string {
	var new_uuid = uuid.NewV1().String()
	var init = comms.NewInitId(new_uuid)
	log.Println("New Client Connected")
	log.Println(init.Id)
	var err = conn.WriteJSON(init)
	if err != nil {
		log.Println(err)
	}
	var newUser = comms.NewUser("", conn)
	newUser.Id = new_uuid
	s.UserMap[new_uuid] = *newUser
	return new_uuid
}

func (s *Server) readerLogin(conn *websocket.Conn,uniqueId string) {
	for (!s.UserMap[uniqueId].Loginned) {
		// read in a message
		var m = comms.InitUser{}
		err := conn.ReadJSON(&m)
		if err != nil {
			log.Println("Error reading json.", err)
		}
		var u = s.UserMap[uniqueId]
		u.UserName = string(m.Name)
		u.Loginned = true
		s.UserMap[uniqueId] = u

		s.DB.InsertIntoUserDB(uniqueId, string(m.Name), "" , "", "")
		log.Println("New user with name: " + string(m.Name))
		
	}
	return
}



// TODO make writing of map with name:key
func (s *Server) writerContacts(conn *websocket.Conn, name string) {
	var users []string
	for _, v := range s.UserMap {
        users = append(users, v.UserName)
    }
	var init = comms.NewUsersOption(users)
	var err = conn.WriteJSON(init)

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

		var key = s.DB.GetUserIdByUsername(string(m.To))

		// log.Println(key)
		err = s.UserMap[key].Conn.WriteJSON(m)
		log.Println("Sended back")
		
		if err != nil {
			log.Println(err)
		}
	}
}


func (s *Server) wsEndpoint(w http.ResponseWriter, r *http.Request) {

	s.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// Init User
	var new_uuid = s.writerNewInitId(ws)
	var u = s.UserMap[new_uuid]
	u.Conn = ws
	s.UserMap[new_uuid] = u

	s.readerLogin(ws, new_uuid)
	
	s.writerContacts(ws, s.UserMap[new_uuid].UserName)

	s.redirectMesasage(ws)

}

func (s *Server) RunServer() {
	http.HandleFunc("/ws", s.wsEndpoint)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.Port), nil))
}
