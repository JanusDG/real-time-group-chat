package handlers

type Hub struct {
}

func (h *Hub) Join() bool {
	return false
}

func (h *Hub) Send() bool {
	return false
}

func (h *Hub) Leave() bool {
	return false
}
