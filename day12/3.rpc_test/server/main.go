package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

/*
	golang写RPC程序，必须符合4个基本条件，不然RPC用不了
	结构体字段首字母要大写，可以别人调用
	函数名必须首字母大写
	函数第一参数是接收参数，第二个参数是返回给客户端的参数，必须是指针类型
	函数还必须有一个返回值error
	练习：模仿前面例题，自己实现RPC程序，服务端接收2个参数，可以做乘法运算，也可以做商和余数的运算，客户端进行传参和访问，得到结果如下：

*/

// Arithmetic 定义用于注册的结构体
type Arithmetic struct {
}

// ArithmeticRequest 定义用于请求的参数
type ArithmeticRequest struct {
	P1, P2 int
}
type ArithmeticResponse struct {
	Ji    int
	Shang int
	YuShu int
}

func (a *Arithmetic) ChengFun(req ArithmeticRequest, ret *ArithmeticResponse) error {
	ret.Ji = req.P1 * req.P2
	return nil
}
func (a *Arithmetic) ChuFun(req ArithmeticRequest, ret *ArithmeticResponse) error {
	if req.P2 == 0 {
		return errors.New("除数不能为0")
	}

	ret.Shang = req.P1 / req.P2
	ret.YuShu = req.P1 % req.P2
	return nil
}

func main() {
	//	1.注册服务
	rect := new(Arithmetic)
	//	注册一个rect服务
	rpc.Register(rect)
	//	2.服务处理绑定到http协议上
	rpc.HandleHTTP()
	//	3.监听服务
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
