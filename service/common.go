package service

import (
	"WebBackend/catch_error"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

//校验密码
func confirmEncryption(kernel string,item string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(kernel), []byte(item))
	return err == nil
}
//制造token
func produceToken(c *gin.Context,id uint)string{
	claims:= struct {
		ID uint `json:"id"`
		jwt.StandardClaims
	}{
		id,
		jwt.StandardClaims{
			ExpiresAt:time.Now().Add(time.Hour * time.Duration(336)).Unix(),
			IssuedAt:time.Now().Unix(),
		},
	}//exp和iat使用科学计数法
	some:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err:=some.SignedString([]byte(os.Getenv("JWT_SECRET")))
	//密文加密失败
	catch_error.CheckWithServer(c,err)
	return token
}

