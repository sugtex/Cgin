package cache

import (
	"github.com/go-redis/redis"
	"os"
)

var Redis *redis.Client
func InitRedis(){
   Redis=redis.NewClient(&redis.Options{
   	Addr:os.Getenv("REDIS_ADDR"),
   	DB:0,
   })
   _,err:=Redis.Ping().Result()
   if err!=nil{
   	panic(err)
   }

}
