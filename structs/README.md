# Structs

Provide some extends util functions for struct. eg: convert, tag parse, struct data init

- `structs.Aliases` - implemented a simple string alias map.
- Convert a struct to `map[string]any` data
- Quickly init struct default values by field "default" tag.
- Quickly set struct field values by map data
- Parse a struct and collect tags, and parse tag value
- And more util functions ...

## Install

```shell
go get github.com/gookit/goutil/structs
```

## Go docs

- [Go docs](https://pkg.go.dev/github.com/gookit/goutil/structs)

## Usage

### Convert to map

`structs.ToMap()` can be quickly convert a `struct` value to `map[string]any`

**Examples**:

```go
	type User1 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		city string
	}

	u1 := &User1{
		Name: "inhere",
		Age:  34,
		city: "somewhere",
	}

	mp := structs.ToMap(u1)
	dump.P(mp)
```

**Output**:

```shell
map[string]interface {} { #len=2
  "age": int(34),
  "name": string("inhere"), #len=6
},
```

### Init default values

`structs.InitDefaults` Quickly init struct default value by field "default" tag.

**Examples**:

```go
type Extra struct {
    City   string `default:"chengdu"`
    Github string `default:"https://github.com/inhere"`
}
type User struct {
    Name  string `default:"inhere"`
    Age   int    `default:"30"`
    Extra Extra
}

u := &User{}
_ = structs.InitDefaults(u, &structs.InitOptions{})
dump.P(u)
```

**Output**:

```shell
&structs_test.User {
  Name: string("inhere"), #len=6
  Age: int(30),
  Extra: structs_test.Extra {
    City: string("chengdu"), #len=7
    Github: string("https://github.com/inhere"), #len=25
  },
},
```

### Set values from map

```go
data := map[string]interface{}{
    "name": "inhere",
    "age":  234,
    "tags": []string{"php", "go"},
    "city": "chengdu",
}

type User struct {
    Name string   `json:"name"`
    Age  int      `json:"age"`
    Tags []string `json:"tags"`
    City string   `json:"city"`
}

u := &User{}
err := structs.SetValues(u, data)
dump.P(u)
```

**Output**:

```shell
&structs_test.User {
  Name: string("inhere"), #len=6
  Age: int(234),
  Tags: []string [ #len=2
    string("php"), #len=3
    string("go"), #len=2
  ],
  City: string("chengdu"), #len=7
},
```

### Tags collect and parse

Parse a struct for collect tags, and parse tag value

**Examples**:

```go
// eg: "desc;required;default;shorts"
type MyCmd struct {
    Name string `flag:"set your name;false;INHERE;n"`
}

c := &MyCmd{}
p := structs.NewTagParser("flag")

sepStr := ";"
defines := []string{"desc", "required", "default", "shorts"}
p.ValueFunc = structs.ParseTagValueDefine(sepStr, defines)

goutil.MustOK(p.Parse(c))
dump.P(p.Tags())
```

Output:

```shell
map[string]maputil.SMap { #len=1
  "Name": maputil.SMap { #len=1
    "flag": string("set your name;false;INHERE;n"), #len=28
  },
},
```

**Parse tag value**

```go
info, _ := p.Info("Name", "flag")
dump.P(info)
```

Output:

```shell
maputil.SMap { #len=4
  "desc": string("set your name"), #len=13
  "required": string("false"), #len=5
  "default": string("INHERE"), #len=6
  "shorts": string("n"), #len=1
},
```

## Testings

```shell
go test -v ./structs/...
```

Test limit by regexp:

```shell
go test -v -run ^TestSetByKeys ./structs/...
```
