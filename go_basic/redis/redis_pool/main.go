package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

func testConnection(ctx context.Context, client *redis.Client) error {
	_, err := client.Ping().Result()
	if err != nil {
		return fmt.Errorf("无法连接Redis: %w", err)
	}
	log.Println("成功连接到Redis服务器")
	log.Println("\n当前连接池状态:")
	stats := client.PoolStats()
	log.Printf("  总连接数: %d", stats.TotalConns)
	log.Printf("  空闲连接: %d", stats.IdleConns)
	log.Printf("  等待次数: %d (连接不足时会发生)", stats.Timeouts)
	log.Printf("  命中次数: %d (成功获取连接次数)", stats.Hits)
	return nil
}

var redisClient *redis.Client

func init() {
	// 这不是真正的"创建连接"，而是创建一个连接管理器（连接池）
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0, // 使用默认DB
		// 连接池配置
		PoolSize:     100,              // 最大连接数
		MinIdleConns: 10,               // 最小空闲连接数
		MaxConnAge:   30 * time.Minute, // 连接最大生存时间
		PoolTimeout:  30 * time.Second, // 获取连接的超时时间
		IdleTimeout:  5 * time.Minute,  // 空闲连接超时时间
	})
	ctx := context.Background()

	if err := testConnection(ctx, redisClient); err != nil {
		log.Fatalf("连接测试失败: %v", err)
	}
}
func main() {
	fmt.Println("<UNK>")
	redisClient.Set("aabbcc", "123", 0)
	defer redisClient.Close()
	fmt.Println(redisClient.Get("aabbcc").Result())
	/*
		当您调用 Set() 或 Get() 时，连接池会：
		从池中取出一个连接
		执行命令
		立即归还连接
		整个过程在毫秒级别完成
		当您打印状态时，连接已经归还
	*/
}
