package comms

import (
	"github.com/satori/go.uuid"
	// "github.com/gorilla/websocket"

)

type DataBase struct {
	Groups []Group 
	Users  []User
}

func NewDatabase() *DataBase {
	return &DataBase {Groups: make([]Group, 0), Users: make([]User, 0)}
}

func (d *DataBase) AddUsersToDataBase(users ...*User) {
	for _, user := range users{
		d.Users = append(d.Users, *user)
	}
}

func (d *DataBase) AddGroupsToDataBase(groups ...*Group) {
	for _, group := range groups{
		d.Groups = append(d.Groups, *group)
	}
}

type Group struct {
	Name string
	Members []User
}

func NewGroup(name string) *Group {
	return &Group{Name: name, Members: make([]User, 0)}
}

func (g *Group) AddToGroup(users ...*User) {
	for _, user := range users{
		g.Members = append(g.Members, *user)
	}
}


// TODO posibly merge User and InitUser??
type User struct {
	Name string
	Loginned bool
	Id uuid.UUID 
}

func NewUser(name string) *User {
	return &User{Name: name, Loginned: false}
}

type InitUser struct {
	Id uuid.UUID 
}

func NewInitUser(id uuid.UUID) *InitUser {
	return &InitUser{Id: id}
}

type Message struct {
	From uuid.UUID 
	To string 
	Content string
}

func NewMessage(from uuid.UUID, to string, content string) *Message {
	return &Message{From: from, To: to, Content: content}
}
