package main

import (
	"fmt"
	// "log"
	// "database/sql"
	_ "github.com/lib/pq"
	// "github.com/satori/go.uuid"

	"github.com/JanusDG/real-time-group-chat/serverside/config"
	"github.com/JanusDG/real-time-group-chat/serverside/server"
	"github.com/JanusDG/real-time-group-chat/serverside/server/database"
	// "github.com/JanusDG/real-time-group-chat/odt"
)


func main() {
	var db = database.NewDatabse()
	defer db.DB.Close()

	// return
	var cfg = config.GetConf()

	serv := server.NewServer(db, cfg.Server.Port,
		cfg.Server.DEBUG_ON)

	if cfg.DEBUG_ON {
		fmt.Println("Created server instance on port:", serv.Port)
	}

	serv.RunServer()
}


// package main

// import (
// 	// "fmt"
// 	// "log"
	

// 	"github.com/JanusDG/real-time-group-chat/serverside/config"
// 	"github.com/JanusDG/real-time-group-chat/serverside/server"
// 	"github.com/JanusDG/real-time-group-chat/serverside/server/database"
// 	// "github.com/JanusDG/real-time-group-chat/odt"
// )


// func main() {
// 	var db = database.InitDatabase()
// 	defer db.Close()
// 	database.InsertIntoUserDB(db, "Arnold", "qwert123")
// 	return
// 	var cfg = config.GetConf()
	
// 	// var initdatabase = InitHardcodeDatabase()

// 	serv := server.NewServer(db, cfg.Server.Port,
// 		cfg.Server.DEBUG_ON)

// 	if cfg.DEBUG_ON {
// 		fmt.Println("Created server instance on port:", serv.Port)
// 	}

// 	serv.RunServer()
// }
