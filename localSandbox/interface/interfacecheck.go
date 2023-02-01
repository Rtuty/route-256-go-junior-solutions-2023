package _interface

import (
	"fmt"
	"reflect"
)

func TestInterf() {
	m := map[interface{}]interface{}{}

	m["one"] = 1
	m["two"] = "2"
	m["three"] = true
	m["four"] = 4.5
	m[1] = 4

	for k, v := range m {
		switch v.(type) {
		case int:
			fmt.Printf("%s is an integer\n", k)
		case string:
			fmt.Printf("%s is an string\n", k)
		case bool:
			fmt.Printf("%s is an boolean\n", k)
		default:
			fmt.Printf("%s is %v\n", k, reflect.TypeOf(v))
		}
	}
	struc := abc{}
	actual := struc.Test2Interf()

	fmt.Print(actual)

}

type abc struct {
	a string `json:"a"`
	b string `json:"b"`
	c string `json:"c"`
}

type def struct {
	d string `json:"d"`
	e string `json:"e"`
	f string `json:"f"`
}

type testInter interface {
	abc
	def
}

func (abc *abc) Test2Interf() error {
	res := map[interface{}]interface{}{}

	res[abc.a] = "a"
	res[abc.b] = "b"
	res[abc.c] = "c"
	res["dfdfd"] = "dfsfsf"

	for k, v := range res {
		fmt.Printf("%s is a key\n", k)
		fmt.Printf("%s is a value\n", v)
		fmt.Println("\n", k)
		fmt.Println("\n", v)
	}

	return nil
}
