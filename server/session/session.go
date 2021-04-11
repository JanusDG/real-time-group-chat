package session

import (
	"fmt"
)

type Connection struct {
	status string
}



type Session struct {
	connection Connection
}

func (session Session) Read() bool {
	return 
}

func (session Session) Write() bool {
	return 
}

func (session Session) Send() bool {
	return 
}