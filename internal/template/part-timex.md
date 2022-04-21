#### Usage

- **Convert time to date by template**

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

