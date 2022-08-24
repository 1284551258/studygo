package taillog

import (
	"fmt"
	"github.com/1284551258/logAgent/etcd"
	"time"
)

var taskMgr *taillogMgr

type taillogMgr struct {
	LogEntry []*etcd.LogEntry
	tskMap   map[string]*TailTask
	confChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	taskMgr = &taillogMgr{
		LogEntry: logEntryConf,
		tskMap:   make(map[string]*TailTask, 16),
		confChan: make(chan []*etcd.LogEntry),
	}
	for _, logEntry := range logEntryConf {

		//每个日志任务开启一个task
		tailTask := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s-%s", tailTask.topic, tailTask.path)
		taskMgr.tskMap[mk] = tailTask
	}
	go taskMgr.run()
}

func NewConfChan() chan<- []*etcd.LogEntry {
	return taskMgr.confChan
}
func (t *taillogMgr) run() {
	for {
		select {
		case newConf := <-t.confChan:
			fmt.Println("新的配置来了：", newConf)
			//1.新增配置
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s-%s", conf.Topic, conf.Path)
				if _, ok := t.tskMap[mk]; !ok {
					//在之前的配置没找到
					fmt.Printf("未找到%s的配置,开始新增", mk)
					tailTask := NewTailTask(conf.Path, conf.Topic)
					taskMgr.tskMap[mk] = tailTask
					fmt.Printf("%s新增结束", mk)
				}
			}

			//2.删除任务，跟原理的LogEntry比较
			for _, entry := range t.LogEntry {
				isDelete := true
				for _, logEntry := range newConf {
					if entry.Topic == logEntry.Topic && entry.Path == logEntry.Path {
						isDelete = false
						break
					}
				}
				if isDelete {
					//开始删除任务
					mk := fmt.Sprintf("%s-%s", entry.Topic, entry.Path)
					t.tskMap[mk].cancel()
				}
			}

		default:
			time.Sleep(time.Second)
		}
	}
}
