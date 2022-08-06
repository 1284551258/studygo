package main

import (
	"fmt"
	"io"
	"os"
)

//借助io.Copy()实现一个拷贝文件函数

func copyFile(dtName,srcName string) (written int64 , err error){
	f, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("打开源文件失败，错误为%v",err)
		return 
	}
	defer f.Close()

	f2, err := os.OpenFile(dtName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开目标文件失败，错误为%v",err)
		return 
	}
	written, err = io.Copy(f2, f)
	if err != nil {
		fmt.Printf("拷贝文件失败，错误为%v",err)
		return 
	}
	return
}

func main4() {
	_, err := copyFile("test2.txt", "test.txt")
	if err != nil {
		fmt.Printf("拷贝失败，错误为：%v",err)
		return 
	}
	fmt.Println("拷贝成功")

}