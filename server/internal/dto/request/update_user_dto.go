package request

type UpdateUserDTO struct {
	Name  string `json:"name"  validate:"omitempty,min=2,max=100"`
	Email string `json:"email" validate:"omitempty,email"`
	Phone string `json:"phone" validate:"omitempty,len=14"`
	Cpf   string `json:"cpf"   validate:"omitempty,len=11,cpf"`
}
