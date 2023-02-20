package jwt

import (
	"net/http"
	"time"

	"github.com/Garfield247/gonf.git/pkg/e"
	"github.com/Garfield247/gonf.git/utils"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var data interface{}

		code = e.StatusOK
		token := ctx.Query("token")
		if token == "" {
			code = e.InvaldToken
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.FaildParseToken
			} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
				code = e.TokenTimeout
			}
		}
		if code != e.StatusOK {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.StatusText(code),
				"data": data,
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
