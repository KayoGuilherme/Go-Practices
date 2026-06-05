package mapper

import (
	"strconv"

	"server/internal/domain"
	"server/internal/dto/response"
	"server/internal/util/pagination"
)

func ToUserResponse(user *domain.User) response.UserResponseDTO {
	if user == nil {
		return response.UserResponseDTO{}
	}
	return response.UserResponseDTO{
		ID:        strconv.FormatInt(user.ID, 10),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []domain.User) []response.UserResponseDTO {
	out := make([]response.UserResponseDTO, 0, len(users))
	for i := range users {
		out = append(out, ToUserResponse(&users[i]))
	}
	return out
}

func ToPaginatedResponse(page *pagination.Response[domain.User]) pagination.Response[response.UserResponseDTO] {
    return pagination.Response[response.UserResponseDTO]{
        Data:    ToUserResponses(page.Data),
        Total:   page.Total,
        Limit:   page.Limit,
        Offset:  page.Offset,
        HasMore: page.HasMore,
    }
}

func ToUserUpdateResponse(user *domain.User, id int) response.UserUpdateResponseDTO {
	if user == nil {
		return response.UserUpdateResponseDTO{}
	}
	userID := user.ID
	if userID == 0 {
		userID = int64(id)
	}
	return response.UserUpdateResponseDTO{
		ID:    strconv.FormatInt(userID, 10),
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Cpf:   user.Cpf,
	}
}
