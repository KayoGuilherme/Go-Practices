package response

import "time"


type ProductResponseDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Stock       string    `json:"stock"`
	Is_on_sale  bool      `json:"is_on_sale"`
	Weight      float64   `json:"weight"`
	Width       float64   `json:"width"`
	Diameter    int32     `json:"diameter"`
	Length      int32     `json:"length"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}