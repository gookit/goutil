#### ArrUtil Usage

**check value**:

```go
arrutil.IntsHas([]int{2, 4, 5}, 2) // True
arrutil.Int64sHas([]int64{2, 4, 5}, 2) // True
arrutil.StringsHas([]string{"a", "b"}, "a") // True

// list and val interface{}
arrutil.Contains(list, val)
arrutil.Contains([]uint32{9, 2, 3}, 9) // True
```

**convert**:

```go
ints, err := arrutil.ToInt64s([]string{"1", "2"}) // ints: []int64{1, 2} 
ss, err := arrutil.ToStrings([]int{1, 2}) // ss: []string{"1", "2"}
```
