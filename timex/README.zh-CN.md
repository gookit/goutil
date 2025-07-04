# Timex

提供增强的时间 `time.Time` 实现，并添加更多常用的函数式方法。

## 安装

```go
go get github.com/gookit/goutil/timex
```

## 开始使用

### 创建timex实例

```go
now := timex.Now()

// from time.Time
tx := timex.New(time.Now())
tx := timex.FromTime(time.Now())

// from time unix
tx := timex.FromUnix(1647411580)
```

Create from datetime string:

```go
// auto match layout by datetime
tx, err  := timex.FromString("2022-04-20 19:40:34")
// custom set the datetime layout
tx, err  := timex.FromString("2022-04-20 19:40:34", "2006-01-02 15:04:05")
// use date template as layout
tx, err  := timex.FromDate("2022-04-20 19:40:34", "Y-m-d H:I:S")
```

### 使用 timex 实例

```go
tx := timex.Now()
```

**Change time**:

```go
tx.Yesterday()
tx.Tomorrow()

tx.DayStart() // get time at Y-m-d 00:00:00
tx.DayEnd() // get time at Y-m-d 23:59:59
tx.HourStart() // get time at Y-m-d H:00:00
tx.HourEnd() // get time at Y-m-d H:59:59

tx.AddDay(2)
tx.AddHour(1)
tx.AddMinutes(15)
tx.AddSeconds(120)
```

**Compare time**:

```go
// before compare
tx.IsBefore(u time.Time)
tx.IsBeforeUnix(1647411580)
// after compare
tx.IsAfter(u time.Time)
tx.IsAfterUnix(1647411580)
```

### Helper functions

```go
ts := timex.NowUnix() // current unix timestamp

t := timex.NowAddDay(1) // from now add 1 day
t := timex.NowAddHour(1) // from now add 1 hour
t := timex.NowAddMinutes(3) // from now add 3 minutes
t := timex.NowAddSeconds(180) // from now add 180 seconds
```

## 按模板将时间转换为日期

**Template Chars**:

```text
 Y,y - year
  Y - year 2006
  y - year 06
 m - month 01-12
 d - day 01-31
 H,h - hour
  H - hour 00-23
  h - hour 01-12
 I,i - minute
  I - minute 00-59
  i - minute 0-59
 S,s - second
  S - second 00-59
  s - second 0-59
... ...
```

| 字符  | 描述                  | 示例输出                        |
|-----|---------------------|-----------------------------|
| `Y` | 4 位数年份              | 2006                        |
| `y` | 2 位数年份              | 06                          |
| `m` | 带前导零的月份 01-12       | 09                          |
| `n` | 没有前导零的月份 1-12       | 9                           |
| `d` | 带前导零的日期  01-31      | 02                          |
| `j` | 没有前导零的日期 1-31       | 2                           |
| `H` | 带前导零的 24 小时格式 00-23 | 15                          |
| `h` | 带前导零的 12 小时格式 01-12 | 3                           |
| `I` | 带前导零的分钟   00-59     | 04                          |
| `i` | 没有前导零的分钟   0-59     | 4                           |
| `S` | 带前导零的第二个   00-59    | 05                          |
| `s` | 没有前导零的第二个  0-59     | 5                           |
| `v` | 毫秒  000             | 530                         |
| `u` | 微秒  000000          | 605080                      |
| `c` | ISO 8601 日期         | `2006-01-02T15:04:05Z07:00` |

> More, please see [charMap](./template.go)

### 使用示例

Examples, use timex format date:

```go
tx := timex.Now()
date := tx.DateFormat("Y-m-d H:I:S") // Output: 2022-04-20 19:09:03
date = tx.DateFormat("y-m-d h:i:s") // Output: 22-04-20 07:9:3
```

**Format time.Time**:

```go
tx := time.Now()
date := timex.DateFormat(tx, "Y-m-d H:I:S") // Output: 2022-04-20 19:40:34
```

More usage:

```go
ts := timex.NowUnix() // current unix timestamp

date := FormatUnix(ts, "2006-01-02 15:04:05") // Get: 2022-04-20 19:40:34
date := FormatUnixByTpl(ts, "Y-m-d H:I:S") // Get: 2022-04-20 19:40:34
```

## ParseRange 时间范围解析

`timex.ParseRange()` 可以简单快速的将相对的时间大小范围、或关键字解析为 `time.Time`

**使用示例**:

```go
start, end, err := ParseRange("-1h~1h", nil)
goutil.PanicErr(err)

fmt.Println(start, end)
```

**支持的表达式格式示例**：

```text
"-5h~-1h"           => 5 hours ago to 1 hour ago
"1h~5h"             => 1 hour after to 5 hours after
"-1h~1h"            => 1 hour ago to 1 hour after
"-1h"               => 1 hour ago to feature. eq "-1h~"
"-1h~0"             => 1 hour ago to now.
"< -1h" OR "~-1h"   => 1 hour ago.
"> 1h" OR "1h"      => 1 hour after to feature

// keyword: now, today, yesterday, tomorrow
"today"          => today start to today end
"yesterday"      => yesterday start to yesterday end
"tomorrow"       => tomorrow start to tomorrow end
```

## Functions

```go
func AddDay(t time.Time, day int) time.Time
func AddHour(t time.Time, hour int) time.Time
func AddMinutes(t time.Time, minutes int) time.Time
func AddSeconds(t time.Time, seconds int) time.Time
func Date(t time.Time, template string) string
func DateFormat(t time.Time, template string) string
func DayEnd(t time.Time) time.Time
func DayStart(t time.Time) time.Time
func Format(t time.Time) string
func FormatBy(t time.Time, layout string) string
func FormatByTpl(t time.Time, template string) string
func FormatUnix(sec int64) string
func FormatUnixBy(sec int64, layout string) string
func FormatUnixByTpl(sec int64, template string) string
func HourEnd(t time.Time) time.Time
func HourStart(t time.Time) time.Time
func HowLongAgo(sec int64) string
func NowAddDay(day int) time.Time
func NowAddHour(hour int) time.Time
func NowAddMinutes(minutes int) time.Time
func NowAddSeconds(seconds int) time.Time
func NowHourEnd() time.Time
func NowHourStart() time.Time
func NowUnix() int64
func SetLocalByName(tzName string) error
func ToDuration(s string) (time.Duration, error)
func ToLayout(template string) string
func TodayEnd() time.Time
func TodayStart() time.Time

// for create timex.Time
    func FromDate(s string, template ...string) (*Time, error)
    func FromString(s string, layouts ...string) (*Time, error)
    func FromTime(t time.Time) *Time
    func FromUnix(sec int64) *Time
    func Local() *Time
    func LocalByName(tzName string) *Time
    func New(t time.Time) *Time
    func Now() *Time
    func Wrap(t time.Time) *Time
```

### Methods in timex.Time

```go
func (t *Time) AddDay(day int) *Time
func (t *Time) AddHour(hours int) *Time
func (t *Time) AddMinutes(minutes int) *Time
func (t *Time) AddSeconds(seconds int) *Time
func (t *Time) CustomHMS(hour, min, sec int) *Time
func (t *Time) DateFormat(template string) string
func (t *Time) Datetime() string
func (t *Time) DayAfter(day int) *Time
func (t *Time) DayAgo(day int) *Time
func (t *Time) DayEnd() *Time
func (t *Time) DayStart() *Time
func (t Time) Diff(u time.Time) time.Duration
func (t Time) DiffSec(u time.Time) int
func (t *Time) Format(layout string) string
func (t *Time) HourEnd() *Time
func (t *Time) HourStart() *Time
func (t Time) HowLongAgo(before time.Time) string
func (t *Time) IsAfter(u time.Time) bool
func (t *Time) IsAfterUnix(ux int64) bool
func (t *Time) IsBefore(u time.Time) bool
func (t *Time) IsBeforeUnix(ux int64) bool
func (t *Time) SubDay(day int) *Time
func (t *Time) SubHour(hours int) *Time
func (t *Time) SubMinutes(minutes int) *Time
func (t *Time) SubSeconds(seconds int) *Time
func (t Time) SubUnix(u time.Time) int
func (t Time) T() time.Time
func (t Time) Timestamp() int64
func (t *Time) Tomorrow() *Time
func (t *Time) TplFormat(template string) string
func (t *Time) UnmarshalJSON(data []byte) error
func (t *Time) UnmarshalText(data []byte) error
func (t *Time) Yesterday() *Time
```

## Code Check & Testing

```bash
gofmt -w -l ./
golint ./...
```

**Testing**:

```shell
go test -v ./timex/...
```

**Test limit by regexp**:

```shell
go test -v -run ^TestSetByKeys ./timex/...
```
