package serializer

import (
	"WebBackend/model"
)

//用户序列化器
type User struct {
	Token string `json:"token"`
	UserName  string `json:"userName"`
	NickName string `json:"nickName"`
	CreatedAt int64  `json:"created_at"`
}
//用户登录序列化
func BuildUserWithToken(u model.User,t string) User{
	return User{
		Token:     t,
		UserName:  u.Account,
		CreatedAt: u.CreatedAt,
	}
}
