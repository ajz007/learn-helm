package main

import (
	"github.com/gin-gonic/gin"
	"go-app1.com/internal/handlers"
	"go.uber.org/zap"
)

// func main() {

// 	// First you create the server Object
// 	router := gin.Default()

// 	// Then you define routes and handlers
// 	router.GET("/hello", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Hello World",
// 		})
// 	})

// 	// You can also set middleware if needed

// 	// Then you start the server in listen mode
// 	log.Println("Starting server on :8080")
// 	router.Run(":8080") // listen and serve on
// }

func main() {

	// Initialize zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// create server object
	serverobj := gin.Default()

	// add routes to the server, handler
	serverobj.GET("/hello", func(c *gin.Context) {
		logger.Info("/hello endpoint called")
		handlers.HelloWorldhandler(c)
	})

	// start the server
	logger.Info("Starting server", zap.String("addr", ":8080"))
	serverobj.Run(":8080") // listen and serve on :8080
}
