package server

import (
		"fmt"
		"log"
		"net/http"
		"strconv"

		"github.com/gorilla/websocket"
)

type Server struct {
	Port int 
	upgrader websocket.Upgrader
	// session session.Session
}


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


//ToDo
// func (s *Server) HandleConnection() bool {
// 	return false
// }

//ToDo
// func (s *Server) CreateSession() bool {
// 	return false
// }

//ToDo
// func (s *Server) RunSession() bool {
// 	return false
// }


