package structs_test

import (
	"fmt"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/structs"
)

func ExampleToMap() {
	type Extra struct {
		City   string `json:"city"`
		Github string `json:"github"`
	}
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Extra Extra  `json:"extra"`
	}

	u := &User{
		Name: "inhere",
		Age:  30,
		Extra: Extra{
			City:   "chengdu",
			Github: "https://github.com/inhere",
		},
	}

	mp := structs.ToMap(u)
	dump.P(mp)
	/*dump:
	map[string]interface {} { #len=3
	  "name": string("inhere"), #len=6
	  "age": int(30),
	  "extra": map[string]interface {} { #len=2
	    "city": string("chengdu"), #len=7
	    "github": string("https://github.com/inhere"), #len=25
	  },
	},
	*/

	fmt.Println("mp.ame:", mp["name"])
	fmt.Println("mp.age:", mp["age"])
	// Output:
	// mp.ame: inhere
	// mp.age: 30
}

func ExampleInitDefaults() {
	type Extra struct {
		City   string `default:"chengdu"`
		Github string `default:"https://github.com/inhere"`
	}
	type User struct {
		Name  string `default:"inhere"`
		Age   int    `default:"30"`
		Extra Extra  `default:""` // add tag for init sub struct
	}

	u := &User{}
	_ = structs.InitDefaults(u)
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

	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
	fmt.Println("Extra.City:", u.Extra.City)
	fmt.Println("Extra.Github:", u.Extra.Github)
	// Output:
	// Name: inhere
	// Age: 30
	// Extra.City: chengdu
	// Extra.Github: https://github.com/inhere
}
