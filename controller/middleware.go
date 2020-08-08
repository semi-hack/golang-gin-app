package controller

import (
	//"fmt"
	//"time"
	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/cors"
)

// SetHeaders ...
func SetHeaders(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}