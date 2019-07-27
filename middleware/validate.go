package middleware

import (
	"WebBackend/cache"
	"WebBackend/catch_error"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"strings"
)
//验证是否是合法客户端发送的请求
func ValidateSign(c *gin.Context){
	h:=c.Request.Header.Get("sign")
	res:=strings.Compare(h,os.Getenv("SIGN"))
	if res==-1{
		catch_error.PanicSelf(c,400,"非法客户端","非法用户"+c.ClientIP())
	}
	c.Next()
}
//验证token
func ValidateToken(c *gin.Context) {
	t:=c.Request.Header.Get("token")
	//没有token
	if t==""{
		catch_error.PanicSelf(c,401,"未登录",c.ClientIP()+"未登录")
	}
	//有token，进行校验
	token, err:= jwt.Parse(t, func(item *jwt.Token)(interface{},error) {
		return []byte(os.Getenv("JWT_SECRET")),nil
	})
	//校验函数出错
	catch_error.CheckWithServer(c,err)
	//解析后是否符合秘钥
	if !token.Valid {
		catch_error.PanicSelf(c,401,"登录权限超时，请重新登陆",c.ClientIP()+"用户权限超时")
	}
	//查询黑名单
	claims:=token.Claims.(jwt.MapClaims)
	key:="b_"+strconv.Itoa(int(claims["id"].(float64)))
	_,err=cache.Redis.Do("get",key).Result()
	if err==redis.Nil{
		c.Next()
	}else if err!=nil{
		catch_error.CheckWithServer(c,err)
	}else{
		catch_error.PanicSelf(c,402,"黑名单成员",c.ClientIP()+"为黑名单成员")
	}
	c.Set("token",token)
	c.Next()
}

