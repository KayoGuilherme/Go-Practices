package handler

import (
	"server/internal/domain"
	"server/internal/dto/request"
	"server/internal/mapper"
	"server/internal/util/http/responses"
	"server/internal/validators"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(authService domain.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var dto request.LoginUserDTO
	if validators.ShouldBindJSON(ctx, &dto) != nil {
		return
	}
	if validators.ValidateStruct(ctx, &dto) != nil {
		return
	}

	token, user, err := h.authService.Login(ctx.Request.Context(), dto.Email, dto.Password)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	ctx.SetCookie("access_token", token, 3600, "/", "localhost", false, true)
	responses.OK(ctx, mapper.ToLoginResponse(token, user))
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "", "", false, true)
	responses.OK(ctx, gin.H{"message": "Logged out successfully"})
}
