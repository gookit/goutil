package structs_test

import (
	"fmt"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/structs"
)

func ExampleInitDefaults() {
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
