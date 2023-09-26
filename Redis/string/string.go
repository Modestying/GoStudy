package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetKey(rdb *redis.Client, key, value string, expiration time.Duration) {
	fmt.Println(rdb.Set(context.Background(), key, value, expiration).Err())
}

func GetKey(rdb *redis.Client, key string) {
	fmt.Println(rdb.Get(context.Background(), key).Result())
}

func IncKey(rdb *redis.Client, key string) {
	fmt.Println(rdb.Incr(context.Background(), key).Result())
}

func DecKey(rdb *redis.Client, key string) {
	fmt.Println(rdb.Decr(context.Background(), key).Result())
}

func TtlKey(rdb *redis.Client, key string) {
	fmt.Println(rdb.TTL(context.Background(), key).Result())
}
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	SetKey(rdb, "StringDemo", "99", time.Second*3)
	GetKey(rdb, "StringDemo")
	IncKey(rdb, "StringDemo")
	GetKey(rdb, "StringDemo")
	DecKey(rdb, "StringDemo")
	TtlKey(rdb, "StringDemo")
	time.Sleep(time.Second * 3)
	
	fmt.Println(rdb.SetNX(context.Background(), "StringDemo", "81", 0).Result())
}
