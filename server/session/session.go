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
	fmt.Println("I'm reading")
	return false
}

func (session Session) Write() bool {
	return false
}

func (session Session) Send() bool {
	return false
}
