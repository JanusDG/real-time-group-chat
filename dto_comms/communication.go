package comms

type InitUser struct {
	Id int 
}

func NewInitUser(id int) *Mess {
	return &Mess{Id: id}
}

type MessageToOther struct {
	From  	int 
	To    	int 
	Content string 
}

func NewMessageToOther(from int, to int, content string) *MessageToOther {
	return &MessageToOther{From: from, To: to, Content: content}
}