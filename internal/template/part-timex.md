#### Examples

**Create timex instance**

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
tx, err  := timex.FromDate("2022-04-20 19:40:34", "Y-M-D H:I:S")
```

**Use timex instance**

```go
tx := timex.Now()
```

change time:

```go
tx.Yesterday()
tx.Tomorrow()

tx.DayStart() // get time at Y-M-D 00:00:00
tx.DayEnd() // get time at Y-M-D 23:59:59
tx.HourStart() // get time at Y-M-D H:00:00
tx.HourEnd() // get time at Y-M-D H:59:59

tx.AddDay(2)
tx.AddHour(1)
tx.AddMinutes(15)
tx.AddSeconds(120)
```

compare time:

```go
// before compare
tx.IsBefore(u time.Time)
tx.IsBeforeUnix(1647411580)
// after compare
tx.IsAfter(u time.Time)
tx.IsAfterUnix(1647411580)
```

**Helper functions**

```go
ts := timex.NowUnix() // current unix timestamp

t := NowAddDay(1) // from now add 1 day
t := NowAddHour(1) // from now add 1 hour
t := NowAddMinutes(3) // from now add 3 minutes
t := NowAddSeconds(180) // from now add 180 seconds
```

**Convert time to date by template**

```text
Template Vars:
 Y,y - year
  Y - year 2006
  y - year 06
 M,m - month 01
 D,d - day 02
 H,h - hour 15
 I,i - minute 04
 S,s - second 05
```

Examples, use timex:

```go
now := timex.Now()
date := now.DateFormat("Y-M-D H:i:s") // Output: 2022-04-20 19:40:34
date = now.DateFormat("y-M-D H:i:s") // Output: 22-04-20 19:40:34
```

Format time.Time:

```go
now := time.Now()
date := timex.DateFormat(now, "Y-M-D H:i:s") // Output: 2022-04-20 19:40:34
```

More usage:

```go
ts := timex.NowUnix() // current unix timestamp

date := FormatUnix(ts, "2006-01-02 15:04:05") // Get: 2022-04-20 19:40:34
date := FormatUnixByTpl(ts, "Y-M-D H:I:S") // Get: 2022-04-20 19:40:34
```
