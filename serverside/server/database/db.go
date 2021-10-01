package database

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"	
	// "github.com/satori/go.uuid"
)



type Database struct {
	DB *sql.DB
}

func NewDatabse() *Database {
	const (
		host     = "127.0.0.1"
		port     = 5432
		user     = "postgres"
		password = "qwerty123"
		dbname   = "Userdata"
	  )

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return &Database{DB: db}
}



func (db *Database) GetUserIdByUsername(username string) string{
	var result string
	var query = fmt.Sprintf("SELECT user_id FROM _user WHERE username = '%s'", username)
	
	if err := db.DB.QueryRow(query).Scan(&result); err != nil {
        log.Fatal(err)
    }
	return result
}

func (db *Database) InsertIntoUserDB(user_id string, username string, name string, surname string, password string) {
	var query = fmt.Sprintf("INSERT INTO _user (user_id, username, name , surname, password) VALUES"+
	"\n('%s','%s','%s','%s','%s')",user_id, username, name , surname, password)
	db.LaunchQuery(query)
}

// func (db *Database) InsertIntoGroupDB(group_id string, group_name string) {
// 	var query = fmt.Sprintf("INSERT INTO _group (group_id, username) VALUES"+
// 	"\n('%s','%s')", group_id, group_name)
// 	db.LaunchQuery(query)
// }

// func (db *Database) AddToGroupDB(group_id string, user_id string ) {
// 	var query = fmt.Sprintf("INSERT INTO _usergroup (group_id, user_id) VALUES"+
// 	"\n('%s','%s')", group_id, user_id)
// 	db.LaunchQuery(query)
// }

func (db *Database) LaunchQuery(query string){
	command, err := db.DB.Query(query)

    if err != nil {
        panic(err.Error())
    }
    defer command.Close()
}
