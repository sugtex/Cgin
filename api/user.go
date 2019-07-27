package api

import (
	"WebBackend/catch_error"
	"WebBackend/model"
	"WebBackend/service"
	"github.com/gin-gonic/gin"
)
//注册接口
func Register(c *gin.Context){
	//表单
	s:=service.UserRegisterService{}
	err:=c.ShouldBind(&s)
	 catch_error.CheckWithClient(c,403,"不符合约束",err)
	s.Register(c)
	successBack(c,nil)
}
//登录
func Login(c *gin.Context) {
	//表单
	s:=service.UserLoginService{}
	err:=c.ShouldBind(&s)
	 catch_error.CheckWithClient(c,403,"不符合约束",err)
	token:=s.Login(c)
	//返回数据
	successBack(c,token)
}
//获得用户问题
func GetQuestion(c *gin.Context){
	account:=c.Query("account")
	length:=len(account)
    if length==0{
		catch_error.PanicSelf(c,450,"账号为空",c.ClientIP()+"未输入账号")
	}else if length<4||length>16{
		catch_error.CheckWithClient(c,403,"不符合约束",nil)
	}
	var user model.User
	err:= model.DB.Where("account = ?", account).First(&user).Error
	catch_error.CheckWithClient(c,451,"账号不存在",err)
	successBack(c,user.Question)
}
//修改密码
func AlterPassword(c *gin.Context){
	var s service.UserAlterPasswordService
	err:=c.ShouldBind(&s)
	 catch_error.CheckWithClient(c,403,"不符合约束",err)
	newToken:=s.AlterPassword(c)
	successBack(c,newToken)
}


