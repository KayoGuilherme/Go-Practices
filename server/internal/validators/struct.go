package validators

import (
	"server/internal/util/http/responses"

	"github.com/gin-gonic/gin"
)

func ValidateStruct(c *gin.Context, obj any) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		responses.UnprocessableEntity(c, FormatErrors(err))
		return err
	}

	if err := Get().Struct(obj); err != nil {
		responses.UnprocessableEntity(c, FormatErrors(err))
		return err
	}

	return nil
}
