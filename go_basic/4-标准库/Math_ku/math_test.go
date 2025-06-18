package Math_ku

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestMath(t *testing.T) {
	a := 1.0
	b := 2.0

	//常量
	maxInt := math.MaxInt
	fmt.Println(maxInt)

	// 比大小
	f := math.Max(a, b)
	fmt.Println(f)

	// 是不是数
	is := math.IsNaN(a)
	fmt.Println(is) //true

	//就是向上取整
	c := 5.6
	ceil := math.Ceil(c)
	fmt.Println(ceil)

	// mod
	mod := math.Mod(5, 3)
	fmt.Println(mod)

	// 平方根
	sqrt := math.Sqrt(9)
	fmt.Println(sqrt)

	//随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //实例化  随机数种子
	fmt.Println(r.Int())                                 // 5577006791947779410
	fmt.Println(r.Int31())                               // 2019727887
	fmt.Println(r.Intn(5))                               // 0
}
