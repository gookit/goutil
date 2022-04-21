#### Usage

- **Convert time to date by template**

```text
Template Vars:
 Y,y - year
 M,m - month
 D,d - month
 H,h - hour
 I,i - minute
 S,s - second
```

Examples, use timex:

```go
now := timex.Now()
date := now.DateFormat("Y-M-D H:i:s") // Output: 2022-04-20 19:40:34
```

Format time.Time:

```go
now := time.Now()
date := timex.DateFormat(now, "Y-M-D H:i:s") // Output: 2022-04-20 19:40:34
```

