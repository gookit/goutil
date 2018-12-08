package utils

import (
	"fmt"
	"github.com/gookit/goutil/json"
	"time"
)

// Go is a basic promise implementation: it wraps calls a function in a goroutine
// and returns a channel which will later return the function's return value.
// from beego/bee
func Go(f func() error) chan error {
	ch := make(chan error)
	go func() {
		ch <- f()
	}()
	return ch
}

// CalcElapsedTime 计算运行时间消耗 单位 ms(毫秒)
func CalcElapsedTime(startTime time.Time) string {
	return fmt.Sprintf("%.3f", time.Since(startTime).Seconds()*1000)
}

// Filling filling a model from submitted data
// data 提交过来的数据结构体
// model 定义表模型的数据结构体
// 相当于是在合并两个结构体(data 必须是 model 的子集)
func Filling(data interface{}, model interface{}) error {
	jsonBytes, _ := json.Encode(data)

	return json.Decode(jsonBytes, model)
}
