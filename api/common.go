package api

import (
	"WebBackend/serializer"
	"github.com/gin-gonic/gin"
)

func successBack(c *gin.Context,data interface{}){
	c.JSON(200,serializer.Response{
		Status:200,
		Data:data,
	})
}