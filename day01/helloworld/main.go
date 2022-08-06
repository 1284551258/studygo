package main

// 导入语句
import "fmt"
// 函数外只能放置标识符、变量、常量、函数、类型的声明

//main没有参数也没有返回值，程序入口
func main0() {
	fmt.Println("Hello world!")
}

/*
编译go build ，需要关闭go111module
1.在项目目录下执行go build  会在当前目录生成helloworld.exe
2.在其他路径执行go build +src后的路径，一直到main.go（github.com\1284551258\day01\helloworl）,会在当前目录生成helloworld.exe
3.指定编译的二进制文件的名称 go build -o hello.exe
4.像脚本文件一样执行go代码，go run main.go
5.go install 先执行编译go build 然后将可执行文件拷到$GOPATH/bin
6.跨平台编译
默认我们go build的可执行文件都是当前操作系统可执行的文件，Go语言支持跨平台编译——在当前平台（例如Windows）下编译其他平台（例如Linux）的可执行文件。

Windows编译Linux可执行文件
如果我想在Windows下编译一个Linux下可执行文件，那需要怎么做呢？只需要在编译时指定目标操作系统的平台和处理器架构即可。

注意：无论你在Windows电脑上使用VsCode编辑器还是Goland编辑器，都要注意你使用的终端类型，因为不同的终端下命令不一样！！！目前的Windows通常默认使用的是PowerShell终端。

如果你的Windows使用的是cmd，那么按如下方式指定环境变量。

SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
如果你的Windows使用的是PowerShell终端，那么设置环境变量的语法为

$ENV:CGO_ENABLED=0
$ENV:GOOS="linux"
$ENV:GOARCH="amd64"
在你的Windows终端下执行完上述命令后，再执行下面的命令，得到的就是能够在Linux平台运行的可执行文件了。

go build
Windows编译Mac可执行文件
Windows下编译Mac平台64位可执行程序：

cmd终端下执行：

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
PowerShell终端下执行：

$ENV:CGO_ENABLED=0
$ENV:GOOS="darwin"
$ENV:GOARCH="amd64"
go build
Mac编译Linux可执行文件
Mac电脑编译得到Linux平台64位可执行程序：

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
Mac编译Windows可执行文件
Mac电脑编译得到Windows平台64位可执行程序：

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
Linux编译Mac可执行文件
Linux平台下编译Mac平台64位可执行程序：

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
Linux编译Windows可执行文件
Linux平台下编译Windows平台64位可执行程序：

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
现在，开启你的Go语言学习之旅吧。人生苦短，let’s Go.

7.go格式化
go fmt main.go
*/