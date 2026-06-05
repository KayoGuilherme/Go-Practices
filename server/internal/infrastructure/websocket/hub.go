package websocket

import "server/internal/domain"

func NewHub() *domain.Hub {
	return &domain.Hub{
		Rooms:      make(map[string]*domain.Room),
		Register:   make(chan *domain.Client),
		Unregister: make(chan *domain.Client),
		Broadcast:  make(chan *domain.Message),
	}
}
