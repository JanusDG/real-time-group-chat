package main

import (
	"fmt"

	"github.com/JanusDG/real-time-group-chat/serverside/config"
	"github.com/JanusDG/real-time-group-chat/serverside/server"
)

func main() {
	var cfg = config.GetConf()

	serv := server.NewServer(cfg.Server.Port,
		cfg.Server.DEBUG_ON)

	if cfg.DEBUG_ON {
		fmt.Println("Created server instance on port:", serv.Port)
	}

	serv.RunServer()
}
