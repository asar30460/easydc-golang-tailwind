package server

type Hub struct {
	Servers    map[int]*Server
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

type Server struct {
	ID       int                 `json:"id"`
	Name     string              `json:"name"`
	Channels map[int]*Channel `json:"channels"`
	Clients  map[int]*Client  `json:"clients"`
}

type Channel struct {
	ID   int
	Name string
}

func NewHub() *Hub {
	return &Hub{
		Servers:   make(map[int]*Server),
		Register:  make(chan *Client),
		Broadcast: make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			// 確保存在client加入Server後，更新Server的clients
			if _, ok := h.Servers[cl.ServerID]; ok {
				s := h.Servers[cl.ServerID]

				if _, ok := s.Clients[cl.ID]; !ok {
					s.Clients[cl.ID] = cl
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Servers[cl.ServerID]; ok {
				if _, ok := h.Servers[cl.ServerID].Clients[cl.ID]; ok {
					if len(h.Servers[cl.ServerID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							ServerID: cl.ServerID,
							Username: cl.Username,
						}
					}

					delete(h.Servers[cl.ServerID].Clients, cl.ID)
					close(cl.Message)
				}
			}
		case msg := <-h.Broadcast:
			if _, ok := h.Servers[msg.ServerID]; ok {
				for _, cl := range h.Servers[msg.ServerID].Clients {
					cl.Message <- msg
				}
			}
		}
	}
}
