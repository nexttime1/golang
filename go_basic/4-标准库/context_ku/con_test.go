package context_ku

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// WithValue
func TestWithValue(t *testing.T) {

	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	key1 := favContextKey("key1")
	ctx := context.WithValue(context.Background(), key1, "Golang")
	f(ctx, key1)
	f(ctx, favContextKey("key2"))
}

func handelrequest(ctx context.Context) {
	go writeredis(ctx)
	go writedatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("handelrequest done.")
			return
		default:
			fmt.Println("handelrequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writeredis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //父进程消失的话
			fmt.Println("writeredis done.")
			return
		default:
			fmt.Println("writeredis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writedatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writedatabase done.")
			return
		default:
			fmt.Println("writedatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}
func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go handelrequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("it's time to stop all sub goroutines!")
	cancel()

	//just for test whether sub goroutines exit or not
	time.Sleep(5 * time.Second)
}

// WithDeadline
const shortDuration = 1 * time.Millisecond

func TestWithDeadline(t *testing.T) {

	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("在截止时间之后停止") //强制停止
	case <-ctx.Done():
		fmt.Println("在截止时间停止") //正常cancel
	}
}

func TestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("在超时时间之后结束")
	case <-ctx.Done():
		fmt.Println("在超时时间结束")
	}
}
