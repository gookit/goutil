package utils

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
	jsonBytes, _ := JsonEncode(data)

	return JsonDecode(jsonBytes, model)
}

// FormatDate
// str eg "2018-01-14T21:45:54+08:00"
func FormatDate(str string) string {
	// 先将时间转换为字符串
	tt, _ := time.Parse("2006-01-02T15:04:05Z07:00", str)

	// 格式化时间
	return tt.Format("2006-01-02 15:04:05")
}

// TransDateToTime
func TransDateToTime(date string) (t time.Time, ok bool) {
	var layout string

	switch len(date) {
	case 10: // 2006-01-02
		layout = "2006-01-02"
	case 19: // 2006-01-02 12:24:36
		layout = "2006-01-02 15:04:05"
	default:
		return
	}

	t, err := time.ParseInLocation(layout, date, time.Local)
	ok = err == nil

	return
}
