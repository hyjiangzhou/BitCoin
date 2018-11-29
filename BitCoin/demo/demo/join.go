package main

import (
	"strings"
	"fmt"
	"bytes"
)

func main()  {
	str1 := []string{"hello", "world", "itcast"}

	res := strings.Join(str1, "=")
	fmt.Printf("res : %s\n", res)


	//bytes.Join
	bytesArray := bytes.Join([][]byte{[]byte("hello"), []byte("world"), []byte("itcast")}, []byte{})
	fmt.Printf("bytesArray : %s", string(bytesArray))
}
