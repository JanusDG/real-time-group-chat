package main

import "github.com/JanusDG/real-time-group-chat/server"

func main() {
	//cfg := NewConfig()

	serv := server.NewServer()

	serv.RunSession()
}
