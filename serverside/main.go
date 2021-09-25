package main

import (
	"fmt"

	"github.com/JanusDG/real-time-group-chat/serverside/config"
	"github.com/JanusDG/real-time-group-chat/serverside/server"
	"github.com/JanusDG/real-time-group-chat/odt"
)

func InitHardcodeDatabase() *comms.DataBase{
	var Arnold = comms.NewUser("Arnold")
	var Alfred = comms.NewUser("Alfred")
	var Anna = comms.NewUser("Anna")
	var Bob = comms.NewUser("Bob")
	var Beatrice = comms.NewUser("Beatrice")

	var A = comms.NewGroup("A")
	A.AddToGroup(Alfred, Anna, Arnold)
	var B = comms.NewGroup("B")
	B.AddToGroup(Bob,Beatrice)

	var D = comms.NewDatabase()
	D.AddUsersToDataBase(Alfred, Anna, Arnold, Bob,Beatrice)
	D.AddGroupsToDataBase(A, B)
	return D
}

func main() {
	var cfg = config.GetConf()

	var initdatabase = InitHardcodeDatabase()

	serv := server.NewServer(initdatabase, cfg.Server.Port,
		cfg.Server.DEBUG_ON)

	if cfg.DEBUG_ON {
		fmt.Println("Created server instance on port:", serv.Port)
	}

	serv.RunServer()
}
