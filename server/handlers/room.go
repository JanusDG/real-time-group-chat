package handlers

type Room struct{}

func (r *Room) Join() bool {
	return false
}

func (r *Room) BroadCast() bool {
	return false
}

func (r *Room) Leave() bool {
	return false
}
