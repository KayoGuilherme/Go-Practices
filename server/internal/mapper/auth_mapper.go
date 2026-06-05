package mapper

import (
	"strconv"

	"server/internal/domain"
	"server/internal/dto/response"
)

func ToLoginResponse(token string, user *domain.User) response.LoginResponseDTO {
	if user == nil {
		return response.LoginResponseDTO{AccessToken: token}
	}
	return response.LoginResponseDTO{
		AccessToken: token,
		ID:          strconv.FormatInt(user.ID, 10),
		Name:        user.Name,
	}
}
