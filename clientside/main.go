package main

import (
	// "fmt"
	"real-time-group-chat/clientside/client"
	// "real-time-group-chat/clientside/connection"
)


func main(){
	// user := client.NewUser(1, "Steven")

	client := client.NewClient()


	client.Init()
}