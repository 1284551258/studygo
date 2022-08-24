package main

import (
	"fmt"
	"github.com/1284551258/logTransfer/es"
	"github.com/1284551258/logTransfer/kafka"

	"github.com/1284551258/logTransfer/conf"
	"gopkg.in/ini.v1"
)

func main() {
	//0.加载配置文件
	//定义一个配置文件变量
	var cfg conf.LogTransferCfg
	err := ini.MapTo(&cfg, "./conf/conf.ini")
	if err != nil {
		fmt.Printf("get config failed,err:%v", err)
		return
	}
	fmt.Printf("cfg:%v", cfg)
	//1.2初始化es
	err = es.Init(cfg.EsCfg.Address, cfg.EsCfg.ChSize, cfg.EsCfg.Nums)
	if err != nil {
		fmt.Printf("es init failed,err:%v", err)
		return
	}
	fmt.Println("es init successful!")
	//1.初始化
	//1.1初始化kafka
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("kafka init failed,err:%v", err)
		return
	}
	fmt.Println("kafka init successful!")
	//2.从kafka消费数据

	//3.发往es
	select {}
}
