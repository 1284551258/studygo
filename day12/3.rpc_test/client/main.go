package main

import (
	"log"
	"net/rpc"
)

type ArithmeticRequest struct {
	P1, P2 int
}
type ArithmeticResponse struct {
	Ji    int
	Shang int
	YuShu int
}

func main() {
	//1.连接到服务端
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}

	//2.调用方法
	req := ArithmeticRequest{
		P1: 10,
		P2: 3,
	}
	ret := ArithmeticResponse{}
	err = conn.Call("Arithmetic.ChengFun", req, &ret)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("ret:%#v", ret)

	ret2 := ArithmeticResponse{}
	err = conn.Call("Arithmetic.ChuFun", req, &ret2)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("ret2:%#v", ret2)
}
