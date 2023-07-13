package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwtutilities "github.com/youssefhmidi/Backend_in_go/JWT_utilities"
	"github.com/youssefhmidi/Backend_in_go/models"
)

func UseTokenVerification(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request.Header.Get("Authorization")
		token := strings.Split(req, " ")

		if len(token) != 2 {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid header"})
			c.Abort()
			return
		}

		isVerified, err := jwtutilities.IsAuthorized(token[1], secret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		if !isVerified {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Not authorized"})
			c.Abort()
			return
		}
		c.Set("Acces_token", token[1])
		c.Next()
	}
}
