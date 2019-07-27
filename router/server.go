package router

import (
	"WebBackend/api"
	"WebBackend/middleware"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	r:=gin.Default()
	r.Use(middleware.Global)
	r.Use(middleware.ValidateSign)//验证合法客户端
	r.POST("/User", api.Register)//注册
	r.POST("/Login", api.Login)//登录
	r.GET("/User",api.GetQuestion)//获得密保问题
	r.Use(middleware.ValidateToken)//检验token
	r.PUT("/User",api.AlterPassword)//修改密码

	r.Run(":8888")
}
