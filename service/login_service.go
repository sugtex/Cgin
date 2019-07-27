package service

import (
	"WebBackend/catch_error"
	"WebBackend/model"
	"github.com/gin-gonic/gin"
)

//用户登录服务表单数据
type UserLoginService struct {
	  Account  string ` form:"userName" binding:"required,min=4,max=16" `
	  PassWord string ` form:"passWord" binding:"required,len=32,alphanum" `
}
//用户登录
func (this *UserLoginService)Login(c *gin.Context) string {
    var user model.User
    err:= model.DB.Where("account = ?", this.Account).First(&user).Error
    catch_error.CheckWithClient(c,450,"账号不存在",err)
    if !confirmEncryption(user.Password,this.PassWord)  {
    	catch_error.PanicSelf(c,451,"密码错误",c.ClientIP()+"输错密码")
	}
     //获得token
     token:=produceToken(c,user.ID)
	 return token
}

