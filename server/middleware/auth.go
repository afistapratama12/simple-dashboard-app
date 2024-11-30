package middleware

import (
	"net/http"
	"simple-dashboard-server/helper"
	"simple-dashboard-server/model"
	"simple-dashboard-server/wrapper"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from header
		header := c.GetHeader("Authorization")

		if header == "" {
			wrapper.ResponseJSONWithMessage(c, http.StatusUnauthorized, "empty header")
			c.Abort()
			return
		}

		headerData := strings.Split(header, " ")

		var claims = &model.Claims{}
		var err error

		if strings.ToLower(headerData[0]) == "bearer" && len(headerData) == 2 {
			claims, err = helper.ValidateToken(headerData[1])
		} else if len(headerData) == 1 {
			claims, err = helper.ValidateToken(headerData[0])
		}

		if err != nil {
			wrapper.ResponseJSONWithMessage(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()

	}
}
