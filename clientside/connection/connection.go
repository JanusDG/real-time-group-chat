package connection

import (
	"github.com/gorilla/websocket"
)

type Connection struct {
	id int
	conn websocket.Conn

}
