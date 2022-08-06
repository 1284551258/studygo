package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

//使用文件操作相关知识，模拟实现linux平台cat命令的功能。

func cat(r *bufio.Reader) {
	for {
		
		b, err := r.ReadBytes('\n')
		if err != nil {
			fmt.Fprintf(os.Stdout, "%s", b)
			break
		}
		fmt.Fprintf(os.Stdout, "%s", b)
	}
}
func main5() {
	flag.Parse()
	if flag.NArg() == 0 { //没有获取到参数，则从标准输入中读取内容
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Printf("cat文件%s失败，错误为%v", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}
