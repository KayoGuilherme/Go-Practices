package handler

import (
	"net/http"

	"server/internal/domain"
	wsinfra "server/internal/infrastructure/websocket"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	Hub *domain.Hub
}

func NewWSHandler(hub *domain.Hub) *WSHandler {
	return &WSHandler{
		Hub: hub,
	}
}

type RoomResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ClientResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type CreateRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *WSHandler) CreateRoom(ctx *gin.Context) {
	var req CreateRoomRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.Hub.Rooms[req.ID] = &domain.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*domain.Client),
	}

	ctx.JSON(http.StatusOK, req)
}

func (h *WSHandler) JoinRoom(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	roomID := ctx.Param("roomID")
	clientID := ctx.Param("clientID")
	username := ctx.Param("username")

	client := &domain.Client{
		Message:  make(chan *domain.Message),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	message := &domain.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	h.Hub.Register <- client
	h.Hub.Broadcast <- message

	go wsinfra.WriteMessage(conn, client)
	wsinfra.ReadMessage(conn, client, h.Hub)
}

func (h *WSHandler) GetRooms(ctx *gin.Context) {
	rooms := make([]RoomResponse, 0)
	for _, room := range h.Hub.Rooms {
		rooms = append(rooms, RoomResponse{
			ID:   room.ID,
			Name: room.Name,
		})
	}

	ctx.JSON(http.StatusOK, rooms)
}

func (h *WSHandler) GetClientsInRoom(ctx *gin.Context) {
	roomID := ctx.Param("roomID")
	clients := make([]ClientResponse, 0)

	if _, ok := h.Hub.Rooms[roomID]; !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	for _, client := range h.Hub.Rooms[roomID].Clients {
		clients = append(clients, ClientResponse{
			ID:       client.ID,
			Username: client.Username,
		})
	}
	ctx.JSON(http.StatusOK, clients)
}
