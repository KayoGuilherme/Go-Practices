package validators

import (
	"server/internal/util/http/responses"

	"github.com/gin-gonic/gin"
)

func ShouldBindJSON(c *gin.Context, obj any) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		responses.BadRequest(c, err)
		return err
	}
	return nil
}