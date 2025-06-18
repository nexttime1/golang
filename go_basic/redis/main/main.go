package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0, // 使用默认DB
	})
	//redisClient.Set("hello", "world", 0)
	defer redisClient.Close()
	fmt.Println("操作成功")
	// 读取值
	result, err := redisClient.Get("hello").Result()
	if err == redis.Nil {
		fmt.Println("ket not exist")
	} else if err != nil {
		log.Panic(err)
	}
	fmt.Println(result)
	redisClient.Do("set", "skw", "love")
	s, err := redisClient.Get("skw").Result()
	if err == redis.Nil {
		fmt.Println("ket not exist")
	}
	fmt.Println(s)
	//do 为一个 interface{}
	do, _ := redisClient.Do("get", "libai").Result()
	fmt.Printf("%v， 类型是 %T\n", do, do)

	//hash  hget  和  hset
	//redisClient.Do("hset", "user001", "name", "xtm01")
	//redisClient.Do("hset", "user001", "age", 20)
	i01, err := redisClient.Do("hget", "user001", "age").Result()
	if err == redis.Nil {
		fmt.Println("ket not exist")
	}
	fmt.Printf("%v， 类型是 %T\n", i01, i01)

	//hmget   hmset
	//redisClient.Do("hmset", "user003", "age", "10", "name", "Cat")
	r1, err := redisClient.Do("hmget", "user003", "name", "age").Result()
	if err == redis.Nil {
		fmt.Println("ket not exist")
	}
	fmt.Printf("%v <类型是> %T\n", r1, r1)
	result01, ok := r1.([]interface{})
	if !ok {
		fmt.Println("不是 []interface {}类型")
		return
	}
	for _, v := range result01 {
		if v == "name" {
			fmt.Println("name ", v)
		} else {
			fmt.Println("age ", v)
		}

	}
}
