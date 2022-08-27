package api

import "github.com/gin-gonic/gin"

func LoadUrls() *gin.Engine {
	// Create a Gin router with default middleware:
	urls := gin.New()
	// Logger middleware will write the logs to gin.DefaultWriter that is os.Stdout even if you set with GIN_MODE=release.
	urls.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	urls.Use(gin.Recovery())
	// The router is a struct that contains a map of handlers.
	// The handlers are functions that take in a context and return a response.
	// The context is a struct that contains a lot of metadata about the request.
	urls.GET("/", rootViews)

	urls.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return urls
}
