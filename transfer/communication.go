package transfer

type InitUser struct {
	Id int 
}

func NewInitUser(id int) *InitUser {
	return &InitUser{Id: id}
}

type Message struct {
	From int 
	To int
	Content string
}

func NewMessage(from int, to int, content string) *Message {
	return &Message{From: from, To: to, Content: content}
}