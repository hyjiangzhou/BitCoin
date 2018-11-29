package main

import "fmt"

type Test struct {
	s string
}

func (t *Test) String() string{
	return fmt.Sprintf("hello world: %s\n", t.s)
}

func main() {
	n := new(Test)
	fmt.Println(n) //hello world:
	//n.s = "你好"
	//d := fmt.Sprintf("%s", n) //
	//fmt.Println(d)
}
