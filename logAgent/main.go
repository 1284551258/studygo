package main

import (
	"fmt"
	"github.com/1284551258/logAgent/conf"
	"github.com/1284551258/logAgent/etcd"
	"github.com/1284551258/logAgent/kafka"
	"github.com/1284551258/logAgent/taillog"
	"github.com/1284551258/logAgent/utils"
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

var cfg = new(conf.AppConf)

func main() {
	//0.加载配置文件
	err := ini.MapTo(cfg, "./conf/conf.ini")
	if err != nil {
		fmt.Println("加载配置文件失败，err:", err)
		return
	}
	fmt.Printf("%#v\n", *cfg)
	// 1.初始化kafka
	err = kafka.Init([]string{cfg.KafkaConf.Addr}, cfg.KafkaConf.MaxSize)
	if err != nil {
		fmt.Println("初始化kafka出错，err:", err)
		return
	}
	fmt.Println("初始化kafka成功")
	// 2.初始化etcd
	etcd.Init(cfg.EtcdConf.Addr, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println("初始化etcd出错，err:", err)
		return
	}
	fmt.Println("初始化etcd成功")
	//2.1从etcd获取配置
	//拿到本机ip
	ip, err := utils.GetOutboundIP()
	if err != nil {
		fmt.Println("GetOutboundIP failed,err:", err)
		return
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ip)
	logEntries, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("get conf from etcd failed!,err:%v", err)
		return
	}
	fmt.Printf("get conf from etcd successful!, conf:%v\n", etcdConfKey)
	for i, entry := range logEntries {
		fmt.Printf("%v,%v\n", i, *entry)
	}
	//3.收集日志发往kafka
	taillog.Init(logEntries)

	//3.1在etcd派一个哨兵取监视日志收集项的变化（有变化及时通知我的logAgent实现热加载配置）

	newConfChan := taillog.NewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan)
	wg.Wait()
	// 3.初始化taillog
	//err = taillog.Init(cfg.TaillogConf.FileName)
	//if err != nil {
	//	fmt.Println("初始化taillong出错，err:", err)
	//	return
	//}
	//fmt.Println("初始化taillong成功")
	//run()

}

//func run() {
//	//	1.读取日志
//	for {
//		select {
//		case line := <-taillog.ReadLog():
//			//	2.发送到kafka
//			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
//			fmt.Println("发送数据为：", line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//
//}
