package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var DB *redis.Client

func Connect() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	err := client.Ping().Err()
	if err != nil {
		fmt.Println("222222222")
		panic(err)
	}
	DB = client

}

func RedisTest() {
	StatusCmd := DB.Set("name", "go", 0)
	fmt.Println(StatusCmd.Result()) //OK <nil>
	fmt.Println(StatusCmd.Val())    //OK
	fmt.Println(StatusCmd.Err())    //<nil>

	stringCmd := DB.Get("name")
	fmt.Println(stringCmd.Val())    //go
	fmt.Println(stringCmd.Result()) //go <nil>
	fmt.Println(stringCmd.Err())    //<nil>
	DB.Set("age", 20, 0)
	stringCmd2 := DB.Get("age")
	fmt.Println(stringCmd2.Int()) //    20 <nil>

	fmt.Println(DB.Exists("name").Val()) //1

	fmt.Println(DB.Incr("age").Val())      //21
	fmt.Println(DB.IncrBy("age", 2).Val()) //23
	fmt.Println(DB.DecrBy("age", 5).Val()) //18

	fmt.Println(DB.TTL("name")) //-1  永不过期
	DB.Set("name", "go", 5*time.Second)
	time.Sleep(5 * time.Second)
	fmt.Println(DB.TTL("name")) // -6

	DB.Set("name1", "宋可为❤薛提猛", 0)
	DB.Expire("name1", 0).Val()
}

// 五大 类型   string list （切片）  hash  （map） set （无序集合）   sorted set  （有序集合）
func main() {
	Connect()
	RedisTest()

}
