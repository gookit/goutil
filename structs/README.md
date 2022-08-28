# Structs

- `structs.Aliases` - implemented a simple string alias map.
- convert struct to map data

## Install

```bash
go get github.com/gookit/goutil/structs
```

## Go docs

- [Go docs](https://pkg.go.dev/github.com/gookit/goutil/structs)

## Usage

### Convert to map

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

### Init defaults

Quickly init struct default value by field "default" tag.

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

	/*dump:
	&structs_test.User {
	  Name: string("inhere"), #len=6
	  Age: int(30),
	  Extra: structs_test.Extra {
	    City: string("chengdu"), #len=7
	    Github: string("https://github.com/inhere"), #len=25
	  },
	},
	*/
```



### Tags collect and parse



## Testings

```shell
go test -v ./structs/...
```

Test limit by regexp:

```shell
go test -v -run ^TestSetByKeys ./structs/...
```
