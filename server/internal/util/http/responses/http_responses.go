package responses

import (
	"errors"
	"net/http"
	"server/internal/domain/exceptions"

	"github.com/gin-gonic/gin"
)

func Created(c *gin.Context, obj any) {
	c.JSON(http.StatusCreated, obj)
}

func OK(c *gin.Context, obj any) {
	c.JSON(http.StatusOK, obj)
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, errorBody(err))
}

func NotFound(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, errorBody(err))
}

func Unauthorized(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, errorBody(err))
}

func Forbidden(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, errorBody(err))
}

func Conflict(c *gin.Context, err error) {
	c.JSON(http.StatusConflict, errorBody(err))
}

func UnprocessableEntity(c *gin.Context, fields map[string]string) {
	c.JSON(http.StatusUnprocessableEntity, fields)
}

func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, errorBody(err))
}

func errorBody(err error) gin.H {
	if err == nil {
		return gin.H{"error": ""}
	}
	return gin.H{"error": err.Error()}
}

func HandleError(c *gin.Context, err error) {
	var appErr *exceptions.AppError

    if errors.As(err, &appErr) {
        c.JSON(appErr.StatusCode, gin.H{"error": appErr.Message})
        return
    }
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
