package middleware

import (
	"github.com/gin-gonic/gin"
	"goMian/config/inner"
	"goMian/router/service"
	"goMian/utils"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("id")
		if id > 0 {
			return
		}
		jwtS, err := c.Cookie("jwt")
		if err != nil {
			service.ErrorMessage(c, err, "not logged in", false)
			return
		}
		claim, err := utils.ParseJWT(jwtS)
		if err != nil {
			service.ErrorMessage(c, err, "", true)
			return
		}
		c.Set("id", claim.UID)
		if time.Now().Add(inner.JWTFlushTime).After(claim.ExpiresAt.Time) {
			//TODO: flush token
		}
	}
}
