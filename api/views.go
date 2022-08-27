package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func rootViews(c *gin.Context) {
	c.JSONP(
		http.StatusOK,
		gin.H{
			"message": "Base api for Go built with Gin",
		},
	)
}
