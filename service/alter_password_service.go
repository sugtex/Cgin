package service

import (
	"WebBackend/cache"
	"WebBackend/catch_error"
	"WebBackend/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type UserAlterPasswordService struct {
	NowPassword string  ` form:"nowPassword" binding:"required,len=32,alphanum"`
	NewPassword string ` form:"newPassword" binding:"required,len=32,alphanum"`
	AgainPassword string` form:"againPassword" binding:"required,len=32,alphanum"`
	Answer string    ` form:"key" binding:"required,len=32,alphanum"`
}
func(this *UserAlterPasswordService)AlterPassword(c *gin.Context)string{
	 token,_:=c.Get("token")//得到验证成功的token
	 claims:=token.(*jwt.Token).Claims.(jwt.MapClaims) //解析出来的类型都是float64
	 id:=int(claims["id"].(float64))
	 tx:=model.DB.Begin()
	 var user model.User
	 //根据解析出来的id得到整个user
	 err:=model.DB.Where("id = ?",id).First(&user).Error
	 catch_error.CheckWithClient(c,450,"账号不存在",err)
	 //确认当前密码
	if !confirmEncryption(user.Password,this.NowPassword){
		catch_error.PanicSelf(c,451,"原始密码错误",c.ClientIP()+"输错原始密码")
	}
	//加密新密码
	hp,err:=bcrypt.GenerateFromPassword([]byte(this.NewPassword), bcrypt.DefaultCost)
	catch_error.CheckWithServer(c,err)
	//更新密码
	 err=model.DB.Model(&user).Update("password",hp).Error
	catch_error.CheckWithClient(c,452,"修改密码失败",err)
	 //得到token的过期时间
	 vtime:=int64(claims["exp"].(float64))//以秒为单位的时间戳
	 //将token存入redis缓存黑名单
	 key:="b_"+strconv.Itoa(id)
	 value:=c.Request.Header.Get("token")
	 _,err=cache.Redis.Do("set",key,value).Result()
	catch_error.CheckWithServer(c,err)
	 _,err=cache.Redis.Do("expireat",key,vtime).Result()
	catch_error.CheckWithServer(c,err)
	 newToken :=produceToken(c,user.ID)
	 tx.Commit()
	 return newToken
}
