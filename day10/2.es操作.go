package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	//初始化client
	client, err := elastic.NewClient(elastic.SetURL("http://11.2.2.128:9200"))
	if err != nil {
		fmt.Printf("connect to es faild,err:%s", err)
		return
	}
	fmt.Println("connect to es successful!")
	s1 := Student{
		Name:    "zhangsan",
		Age:     18,
		Married: false,
	}
	response, err := client.Index().Index("user").BodyJson(s1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s,type %s\n", response.Id, response.Index, response.Type)
}
