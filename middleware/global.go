package middleware

import "github.com/gin-gonic/gin"

func Global(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST,GET,PUT,OPTIONS")
	c.Header("Access-Control-Allow-Headers","*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}

