package handler

import (
	"server/internal/domain"
	"server/internal/dto/request"
	"server/internal/mapper"
	"server/internal/util/pagination"
	"server/internal/util/http/responses"
	"server/internal/validators"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService domain.ProductService
}

func NewProductHandler(svc domain.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: svc,
	}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var dto request.CreateProductDTO
	if validators.ShouldBindJSON(ctx, &dto) != nil {
		return
	}
	if validators.ValidateStruct(ctx, &dto) != nil {
		return
	}

	product := mapper.ToProductDomain(dto)

	product, err := h.productService.CreateProduct(ctx.Request.Context(), product)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.Created(ctx, mapper.ToProductResponse(product))
}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	var params pagination.Params

	if err := validators.ValidateStruct(ctx, &params); err != nil {
		responses.BadRequest(ctx, err)
		return
	}

	products, err := h.productService.GetAllProducts(ctx, params)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.OK(ctx, mapper.ToProductPaginateResponse(products))
}

func (h *ProductHandler) GetProduct(ctx *gin.Context) {
	id, err := validators.ParsePathInt(ctx, "id")
	if err != nil {
		return
	}

	product, err := h.productService.GetProductById(ctx, id)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.OK(ctx, mapper.ToProductResponse(product))
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	var dto request.UpdateProductDTO

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

	product, err := h.productService.UpdateByProduct(ctx.Request.Context(), id, mapper.ToProductDomainUpdate(dto))
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.OK(ctx, mapper.ToProductResponse(product))
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id, err := validators.ParsePathInt(ctx, "id")
	if err != nil {
		return
	}

	err = h.productService.DeleteByProduct(ctx.Request.Context(), id)
	if err != nil {
		responses.HandleError(ctx, err)
		return
	}

	responses.NoContent(ctx)
}
