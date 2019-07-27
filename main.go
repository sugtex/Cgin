package main

import (
	"WebBackend/cache"
	"WebBackend/model"
	"WebBackend/router"
	"github.com/joho/godotenv"
)

func main() {
	//读取配置文件
	godotenv.Load()

	cache.InitRedis()
	model.InitDataBase()
    defer model.DB.Close()
	defer cache.Redis.Close()
	router.RouterInit()

}
