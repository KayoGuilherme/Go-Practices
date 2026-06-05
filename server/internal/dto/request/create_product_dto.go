package request

type CreateProductDTO struct {
	Name        string  `json:"name"         validate:"required,min=2,max=200"`
	Price       float64 `json:"price"        validate:"required,gte=5"`
	Description string  `json:"description"  validate:"required,min=1"`
	Stock       string  `json:"stock"        validate:"required"`
	Is_on_sale  bool    `json:"is_on_sale"`	 	
	Weight      float64 `json:"weight"       validate:"gte=0"`
	Width       float64 `json:"width"        validate:"gte=0"`
	Diameter    int32   `json:"diameter"     validate:"gte=0"`
	Length      int32   `json:"length"       validate:"gte=0"`
}
