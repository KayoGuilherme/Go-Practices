package request

type CreateUserDTO struct {
	Name     string `json:"name"     validate:"required,min=2,max=100"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=72"`
	Phone    string `json:"phone"    validate:"required,len=14"`
	Cpf      string `json:"cpf"      validate:"required,len=11,cpf"`
}