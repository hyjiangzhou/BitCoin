package main

import (
	"encoding/gob"
	"bytes"
	"log"
	"fmt"
)

type Person struct {
	Name string
	Age  uint64
}


func main()  {

	//定义个Person结构，进行编码后传输
	lily := Person{"lily", 28}


	var buffer bytes.Buffer

	//1. 定义编码器
	encoder := gob.NewEncoder(&buffer)
	//2. 使用编码器编码

	err := encoder.Encode(&lily)
	// 一定要记得校验
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("编码后的数据: %x\n", buffer.Bytes())


	//================ 在对端，进行解码
	//对端解码时使用容器
	var p1 Person

	//1. 定义一个解码器
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))

	//var buffer1 bytes.Buffer
	//buffer1.Write(buffer.Bytes())
	//decoder := gob.NewDecoder(&buffer1)

	//2. 对传过来字节流进行解码
	err = decoder.Decode(&p1)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("p1 : %v\n", p1)
}
