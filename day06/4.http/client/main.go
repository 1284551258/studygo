package main

import "fmt"

type tester interface {
	hello()
}

type test struct {
	name string
}

func (t test) hello() {
	fmt.Println("hello")
}

func main() {

	var t1 tester

	fmt.Printf("%T,%v", t1, t1)

	t2 := test{
		name: "张三",
	}
	t1 = t2
	fmt.Printf("%T,%v", t1, t1)
	t1.hello()
	fmt.Printf("%v,%T,%v", t2.name, t1, t1)

}
