package domain

type Client struct {
	Message  chan *Message
	ID       string
	RoomID   string
	Username string
}
