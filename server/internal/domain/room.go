package domain

type Room struct {
	ID      string
	Name    string
	Clients map[string]*Client
}
