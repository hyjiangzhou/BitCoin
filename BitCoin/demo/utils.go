package main

import "os"
import "fmt"
 //判断文件是否存在
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func Welcome() {
	fmt.Printf("\n======================================================================================\n")
	fmt.Printf("                   欢迎参考 江洲 为您提供的开源 bitcoin !\n")
	fmt.Printf(" 更多内容参考：CSDN博客：cqu_jiangzhou ：https://me.csdn.net/weixin_42117918\n")
	fmt.Printf("======================================================================================\n")
}
