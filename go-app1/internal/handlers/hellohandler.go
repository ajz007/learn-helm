package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func HelloWorldhandler(c *gin.Context) {
	logger.Info("HelloWorldhandler called", zap.String("endpoint", "/hello"))
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
