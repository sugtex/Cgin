package service

import (
	"WebBackend/catch_error"
	"WebBackend/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//用户注册服务表单数据
type UserRegisterService struct {
	Account  string ` form:"userName" binding:"required,min=4,max=16"`
	PassWord string `  form:"passWord" binding:"required,len=32,alphanum"`
	Question string  `  form:"matter" binding:"required"  `
	Answer   string  `  form:"key" binding:"required,len=32,alphanum"  `
}
//用户注册
func (this *UserRegisterService)Register(c *gin.Context){
	var user model.User
	//用户名
	user.Account=this.Account
	//创建时间
	user.CreatedAt=time.Now().Unix()
    //二次加密密码
	hp,err:=bcrypt.GenerateFromPassword([]byte(this.PassWord), bcrypt.DefaultCost)
	catch_error.CheckWithServer(c,err)
	user.Password=string(hp)
	//随机产生昵称
	user.NickName=func() string{
		rand.Seed(time.Now().Unix())
		var builder strings.Builder
		builder.Write([]byte("哲学用户_"))
		for i:=0;i<5;i++{
			r:=rand.Intn(100)
			builder.Write([]byte(strconv.Itoa(r)))
		}
		return builder.String()
	}()
	//问题
	user.Question=this.Question
	//二次加密答案
	ha,err:=bcrypt.GenerateFromPassword([]byte(this.Answer), bcrypt.DefaultCost)
	catch_error.CheckWithServer(c,err)
	user.Answer=string(ha)
	//创建用户
	err= model.DB.Create(&user).Error
	catch_error.CheckWithClient(c,450,"用户名重复",err)
}