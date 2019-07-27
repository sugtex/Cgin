package catch_error

import (
	"WebBackend/serializer"
	"github.com/gin-gonic/gin"
)
//服务端出错
func CheckWithServer(c *gin.Context,err error){
	if err != nil {
		c.JSON( 200, serializer.Response{
			Status:500,
			Msg:"服务端出错",
			Error:err.Error(),
		})
		panic(err)
	}
}
//客户端出错
func CheckWithClient(c *gin.Context,code int,msg string,err error){
	if err != nil {
		c.JSON( 200, serializer.Response{
			Status:code,
			Msg:msg,
			Error:err.Error(),
		})
		panic(err)
	}
}
//自己panic
func PanicSelf(c *gin.Context,code int,msg string,err string){
	c.JSON( 200, serializer.Response{
		Status:code,
		Msg:msg,
	})
	panic(err)
}
