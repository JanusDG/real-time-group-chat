package main

import (
		"fmt"
		"os"

		"github.com/JanusDG/real-time-group-chat/server"
		"github.com/JanusDG/real-time-group-chat/config"
)

func main() {
	var config = config.GetConf()

	serv := new(server.Server)
	serv.Init(config.Server.Port,
			  config.Server.DEBUG_ON)

	
	if (config.DEBUG_ON){ fmt.Println("Created server instance on port:",serv.Port) }
	if (config.DEBUG_ON){ fmt.Println(os.Getenv("DEBUG_ON")) }
	
	serv.RunServer()
}
