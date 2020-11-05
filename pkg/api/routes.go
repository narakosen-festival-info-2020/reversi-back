package api

import (
	"net/http"

	"github.com/narakosen-festival-info-2020/reversi-back/pkg/reversi"

	"github.com/gin-gonic/gin"
)

func pingRoute(engine *gin.Engine, server *Info) {
	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, server.GetJSON())
	})
}

func reversiRoute(engine *gin.Engine, server *Info) {
	reveriGroup := engine.Group("/reversi")
	reveriGroup.Use(server.tokenCheck())

	extract := func(ctx *gin.Context) string {
		return ctx.GetHeader("Authorization")[7:]
	}

	reveriGroup.GET("/state", func(ctx *gin.Context) {
		specificCode := extract(ctx) // before check by Middleware
		data, _ := server.getMatchData(specificCode)
		ctx.JSON(http.StatusOK, data.GetJSON())
	})

	type placeStone struct {
		Y int `json:"y"`
		X int `json:"x"`
	}

	reveriGroup.POST("/state/action", func(ctx *gin.Context) {
		specificCode := extract(ctx)
		var action placeStone
		if err := ctx.ShouldBindJSON(&action); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
			})
			return
		}
		data, _ := server.getMatchData(specificCode)
		response := data.PlaceStone(action.Y, action.X, 1, true)
		if response == -1 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Cannot this action",
			})
			return
		}
		ctx.JSON(http.StatusOK, data.GetJSON())
	})
}

func generateMatchRoute(engine *gin.Engine, server *Info) {
	engine.POST("/generate", func(ctx *gin.Context) {
		var req reversi.GenerateData
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
			})
			return
		}
		check, err := server.generateMatch(&req)
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
