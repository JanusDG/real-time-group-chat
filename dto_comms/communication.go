package comms

type Mess struct {
	Id int 
}

func NewMess(id int) *Mess {
	return &Mess{Id: id}
}