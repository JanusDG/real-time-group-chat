package server

import ( "fmt"
	// "html"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/websocket"
    // "sync"
	// "github.com/JanusDG/real-time-group-chat/server/session"
)


func Dudka() {
	fmt.Println("Piupiu")
}

type Server struct {
	Port int 
	upgrader websocket.Upgrader
	// session session.Session
}

// initializer
// func NewServer(port int) *Server { return &Server{port} }

// constructor
func (s *Server) Init(port int) { 
	s.Port = port 
	
}

func (s *Server) RunServer() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	
	http.HandleFunc(
		"/hello", 
		func (w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "hello")} )


    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(s.Port), nil))
}

// func (s *Server) HandleConnection() bool {
// 	return false
// }

// func (s *Server) CreateSession() bool {
// 	return false
// }

// func (s *Server) RunSession() bool {
// 	return false
// }


