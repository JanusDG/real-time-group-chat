package main

import (
	"github.com/JanusDG/real-time-group-chat/clientside/client"
)

func main(){
	client := client.NewClient()

	client.Init()
}