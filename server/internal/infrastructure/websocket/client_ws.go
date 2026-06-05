package websocket

import (
	"log"

	"server/internal/domain"

	"github.com/gorilla/websocket"
)

func WriteMessage(conn *websocket.Conn, client *domain.Client) {
	defer conn.Close()

	for {
		message, ok := <-client.Message
		if !ok {
			return
		}

		conn.WriteJSON(message)
	}
}

func ReadMessage(conn *websocket.Conn, client *domain.Client, hub *domain.Hub) {
	defer func() {
		hub.Unregister <- client
		conn.Close()
	}()

	for {
		_, m, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		hub.Broadcast <- &domain.Message{
			Content:  string(m),
			RoomID:   client.RoomID,
			Username: client.Username,
		}
	}
}
