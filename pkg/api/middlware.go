package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (info *Info) tokenCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		specificCode := ctx.GetHeader("Authorization")
		if specificCode[0:6] != "Bearer" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
			})
			ctx.Abort()
			return
		}
		specificCode = specificCode[7:]
		_, check := info.getMatchData(specificCode)
		fmt.Println(specificCode)
		if !check {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "Forbidden",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
