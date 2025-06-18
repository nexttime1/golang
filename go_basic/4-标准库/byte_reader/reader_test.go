package byte_reader

import (
	"bytes"
	"fmt"
	"testing"
)

/*
  - func NewReader(b []byte) *Reader`
    将 b 包装成 bytes.Reader 对象。
  - `func (r *Reader) Len() int`
    返回未读取部分的数据长度
  - `func (r *Reader) Size() int64`
    返回底层数据的总长度，方便 ReadAt 使用，返回值永远不变。
  - `func (r *Reader) Reset(b []byte)`
    将底层数据切换为 b，同时复位所有标记（读取位置等信息）。
*/
func TestReader(test *testing.T) {

	reader := bytes.NewReader([]byte("hello world"))

	i := reader.Len()
	fmt.Println("len 为： ", i) //11

	reader.Seek(2, 0)
	buf := make([]byte, 2)
	_, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf)) //ll
	k := reader.Size()
	fmt.Println("Size 为：  ", k) //永远不变 11
	i = reader.Len()
	fmt.Println("len 为： ", i) // 7
}
