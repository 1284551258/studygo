package day12

/*
1.	RPC简介
	远程过程调用（Remote Procedure Call，RPC）是一个计算机通信协议
	该协议允许运行于一台计算机的程序调用另一台计算机的子程序，而程序员无需额外地为这个交互作用编程
	如果涉及的软件采用面向对象编程，那么远程过程调用亦可称作远程调用或远程方法调用

3.	golang中如何实现RPC
	golang中实现RPC非常简单，官方提供了封装好的库，还有一些第三方的库
	golang官方的net/rpc库使用encoding/gob进行编解码，支持tcp和http数据传输方式，由于其他语言不支持gob编解码方式，
所以golang的RPC只支持golang开发的服务器与客户端之间的交互
	官方还提供了net/rpc/jsonrpc库实现RPC方法，jsonrpc采用JSON进行数据编解码，因而支持跨语言调用，
目前jsonrpc库是基于tcp协议实现的，暂不支持http传输方式
	例题：golang实现RPC程序，实现求矩形面积和周长
	服务端
*/
