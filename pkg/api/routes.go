package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingRoute(engine *gin.Engine, server *Info) {
	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, server.GetJSON())
	})
}

func reversiRoute(engine *gin.Engine, server *Info) {
	reveriGroup := engine.Group("/api/reversi")
	reveriGroup.Use(server.tokenCheck())
	reveriGroup.GET("/state", func(ctx *gin.Context) {
		specificCode := ctx.GetHeader("Authorization")[7:]
		data, tmp := server.getMatchData(specificCode)
		fmt.Println(specificCode)
		fmt.Println(tmp)
		ctx.JSON(http.StatusOK, data.GetJSON())
	})
}

func generateMatchRoute(engine *gin.Engine, server *Info) {
	type request struct {
		BoardType string `json:"board_type"`
	}
	engine.POST("/api/generate", func(ctx *gin.Context) {
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
			})
			return
		}
		check, err := server.generateMatch(req.BoardType)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, check.GetJSON())
	})
}

func setRoute(engine *gin.Engine, server *Info) {
	pingRoute(engine, server)
	reversiRoute(engine, server)
	generateMatchRoute(engine, server)
}
