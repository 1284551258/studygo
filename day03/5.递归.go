package main

import "fmt"

//递归，函数自己调用自己 地址：https://www.runoob.com/go/go-recursion.html
//适合处理那种问题相同，问题的规律越来越小的场景
//计算n的阶乘 n!= n*(n-1)*(n-2)*...*3*2*1

func f(n uint64 ) uint64{
	if n > 1{
		ret := n*f(n-1)
		return ret
	}
	return 1
}
//n个台阶，一次可以走1步，也可以走2步，有多少种走法
func taijie(n int) int{
	if n == 1{
		return 1
	} else if n == 2 {
		return 2
	} else {
		return taijie(n-1)+taijie(n-2)
	}
}
func main(){
	fmt.Println("4的阶乘：",f(4))
	fmt.Println("5层台阶一共的走法为：",taijie(5))
}