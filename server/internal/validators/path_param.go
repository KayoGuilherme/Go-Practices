package validators

import (
	"errors"
	"strconv"

	"server/internal/util/http/responses"

	"github.com/gin-gonic/gin"
)

func ParsePathInt(c *gin.Context, param string) (int, error) {
	id, err := strconv.Atoi(c.Param(param))
	if err != nil {
		responses.BadRequest(c, errors.New("invalid id"))
		return 0, err
	}
	return id, nil
}
