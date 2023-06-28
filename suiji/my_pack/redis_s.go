package my_pack

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Do_redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       3,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	x := rdb.Exists("cs")
	x1 := rdb.Exists("cs2")
	fmt.Println(x, x1.Val())
	fmt.Println(rdb.Get("cs").Val())
	fmt.Println(rdb.Get("cs2").Val())
}
