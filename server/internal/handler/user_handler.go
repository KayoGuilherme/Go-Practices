package handler

import (
	"server/internal/domain"
	"server/internal/dto/request"
	"server/internal/mapper"
	"server/internal/util/pagination"
	"server/internal/util/http/responses"
	"server/internal/validators"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	userService domain.UserService
	logger      *zap.Logger
}

func NewUserHandler(userService domain.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var dto request.CreateUserDTO
	if validators.ShouldBindJSON(ctx, &dto) != nil {
		return
	}

	if validators.ValidateStruct(ctx, &dto) != nil {
		return
	}

	user, err := h.userService.CreateUser(ctx.Request.Context(), dto.Name, dto.Email, dto.Password, dto.Phone, dto.Cpf)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.Created(ctx, mapper.ToUserResponse(user))
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	var params pagination.Params

	if err := validators.ValidateStruct(ctx, &params); err != nil {
		h.logger.Warn("parâmetros de paginação inválidos",
			zap.Error(err),
			zap.String("path", ctx.Request.URL.Path),
		)
		responses.BadRequest(ctx, err)
		return
	}

	params.SetDefaults()

	users, err := h.userService.GetAllUsers(ctx.Request.Context(), params)
	if err != nil {
		h.logger.Error("falha ao listar usuários",
			zap.Error(err),
			zap.Int("limit", params.Limit),
			zap.Int("offset", params.Offset),
		)
		responses.HandleError(ctx, err)
		return
	}

	responses.OK(ctx, mapper.ToPaginatedResponse(users))
}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	id, err := validators.ParsePathInt(ctx, "id")
	if err != nil {
		return
	}

	user, err := h.userService.GetUserById(ctx.Request.Context(), id)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.OK(ctx, mapper.ToUserResponse(user))
}

func (h *UserHandler) GetUserByCpf(ctx *gin.Context) {
	user, err := h.userService.GetUserByCpf(ctx.Request.Context(), ctx.Param("cpf"))
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.OK(ctx, mapper.ToUserResponse(user))
}

func (h *UserHandler) DeleteUserById(ctx *gin.Context) {
	id, err := validators.ParsePathInt(ctx, "id")
	if err != nil {
		return
	}

	err = h.userService.DeleteUserById(ctx.Request.Context(), id)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.NoContent(ctx)
}

func (h *UserHandler) UpdateUserById(ctx *gin.Context) {
	var dto request.UpdateUserDTO
	if validators.ShouldBindJSON(ctx, &dto) != nil {
		return
	}
	if validators.ValidateStruct(ctx, &dto) != nil {
		return
	}

	id, err := validators.ParsePathInt(ctx, "id")
	if err != nil {
		return
	}

	user, err := h.userService.UpdateUserById(ctx.Request.Context(), id, dto.Name, dto.Email, dto.Phone, dto.Cpf)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.OK(ctx, mapper.ToUserUpdateResponse(user, id))
}
