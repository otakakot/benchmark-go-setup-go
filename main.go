package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	engine := gin.Default()

	engine.GET("/", func(ctx *gin.Context) {
		logger.Info("handling request", zap.String("path", "/"))

		ctx.JSON(http.StatusOK, "ok")
	})

	engine.GET("/ping", func(ctx *gin.Context) {
		logger.Info("handling request", zap.String("path", "/ping"))

		ctx.JSON(http.StatusOK, "pong")
	})

	engine.Run(":8080")
}
