package comms

import (
	"github.com/gorilla/websocket"

)

// TODO posibly merge User and InitUser??
type User struct {
	UserName string   `json:"username"`
	Loginned bool 
	Conn     *websocket.Conn
	Id string
}

func NewUser(name string, conn *websocket.Conn) *User {
	return &User{UserName: name, Conn: conn, Loginned: false}
}

type InitId struct {
	Id string
}

func NewInitId(id string) *InitId {
	return &InitId{Id: id}
}

type InitUser struct {
	Name string
	Password string
}

func NewInitUser(name string, password string) *InitUser {
	return &InitUser{Name: name, Password: password}
}

type UsersOption struct {
	Users []string
}



func NewUsersOption(users []string) *UsersOption {
	return &UsersOption{Users: users}
}

type Message struct {
	From string
	To string 
	Content string
}

func NewMessage(from string, to string, content string) *Message {
	return &Message{From: from, To: to, Content: content}
}
