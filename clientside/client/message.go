package client

type Message struct{
	From_id int
	To_id int
	Message string
}

func NewMessage(from int, to int, message string) *Message {
	return &Message{From_id: from,
				To_id: to,
				Message: message}
}