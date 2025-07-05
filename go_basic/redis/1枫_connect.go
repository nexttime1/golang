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

func RedisList() {
	DB.LPush("list", "xtm", "libai", "zhaoyun")
	DB.RPush("list", "meng", "sisi", "tutu")
	fmt.Println(DB.LRange("list", 0, -1).Val()) //[zhaoyun libai xtm meng sisi tutu]
	fmt.Println(DB.LPop("list").Val())          //zhaoyun
	fmt.Println(DB.RPop("list").Val())          //tutu
	DB.RPush("list", "aiqiyi")

}
func RedisHash() {
	DB.HSet("hash1", "name", "xtm")
	DB.HSet("hash1", "age", 23)
	fmt.Println(DB.HGet("hash1", "name").Val()) //xtm
	fmt.Println(DB.HGetAll("hash1").Val())      //map[address:2334 age:23 name:xtm]
	fmt.Println(DB.HKeys("hash1").Val())        //[name age address]
	fmt.Println(DB.HLen("hash1").Val())         // 3
}

func RedisSet() {
	DB.SAdd("set1", "1", "2", "3", "4")
	fmt.Println(DB.SMembers("set1").Val())       //[1 2 3 4]
	fmt.Println(DB.SMembersMap("set1").Val())    //map[1:{} 2:{} 3:{} 4:{}]
	fmt.Println(DB.SIsMember("set1", "1").Val()) //true
	fmt.Println(DB.SIsMember("set1", "8").Val()) //false
	fmt.Println(DB.SCard("set1").Val())          //4          个数
	DB.SAdd("SET001", 1, 2, 3, 4, 5)
	DB.SAdd("SET002", 2, 3, 4, 7, 9)
	fmt.Println(DB.SDiff("SET001", "SET002").Val())  //[1 5]
	fmt.Println(DB.SDiff("SET002", "SET001").Val())  //[7 9]
	fmt.Println(DB.SInter("SET001", "SET002").Val()) // [2 3 4]   差集的补
	fmt.Println(DB.SUnion("SET001", "SET002").Val()) // [1 2 3 4 5 7 9]

	fmt.Println(DB.SPop("SET001").Val())
	fmt.Println(DB.SPopN("SET001", 2).Val())

}
func RedisSortedSet() {
	DB.ZAdd("z01", redis.Z{Score: 100, Member: "xtm"},
		redis.Z{Score: 60, Member: "ss"},
		redis.Z{Score: 70, Member: "rr"},
		redis.Z{Score: 101, Member: "skw"},
		redis.Z{Score: 90, Member: "zhao"})

	fmt.Println(DB.ZRange("z01", 0, -1).Val())                                           //[ss rr zhao xtm skw]
	fmt.Println(DB.ZRevRange("z01", 0, -1).Val())                                        //[skw xtm zhao rr ss]
	fmt.Println(DB.ZRangeByScore("z01", redis.ZRangeBy{Min: "80", Max: "110"}).Val())    //[zhao xtm skw]
	fmt.Println(DB.ZRevRangeByScore("z01", redis.ZRangeBy{Min: "80", Max: "110"}).Val()) //[skw xtm zhao]
	fmt.Println(DB.ZRangeWithScores("z01", 0, -1).Val())                                 //[{60 ss} {70 rr} {90 zhao} {100 xtm} {101 skw}]

}
func RedisPipeLine() {
	tx := DB.Pipeline()
	tx.Set("nammm1", "xtm", 0)
	res := tx.Get("nammm1")
	fmt.Println(res.Val()) //没有东西  因为没有提交
	_, err := tx.Exec()
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Val()) //xtm
}
func RedisMutex() {
	DB.Watch(func(tx *redis.Tx) error {
		_, err := tx.Pipelined(func(pipe redis.Pipeliner) error {
			time.Sleep(time.Second * 5)
			pipe.Set("price", 100, 0)
			return nil
		})
		if err != nil {
			fmt.Println("事务更新失败")
			return err
		}
		res := tx.Get("price")
		fmt.Println(res.Val())
		return nil
	}, "price")

}

// 五大 类型   string list （切片）  hash  （map） set （无序集合）   sorted set  （有序集合）
func main() {
	Connect()
	//RedisTest() //string
	//RedisList() //list
	//RedisHash() //hash
	//RedisSet() //set
	//RedisSortedSet() //有序集合
	//RedisPipeLine() //事务
	RedisMutex() //锁
}
