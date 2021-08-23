package main

import ("fmt"
		"os"
		"github.com/JanusDG/real-time-group-chat/server"
		"github.com/JanusDG/real-time-group-chat/config"
)

func main() {
	server.Dudka()

	var config = config.GetConf()

	// serv := server.NewServer(config.Server.Port) 

	serv := new(server.Server)
	serv.Init(config.Server.Port)


	if (config.DEBUG_ON){ fmt.Println("Created server instance on port:",serv.Port) }
	if (config.DEBUG_ON){ fmt.Println(os.Getenv("DEBUG_ON")) }
	



	// server.RunServer()
	serv.RunServer()

	

}
