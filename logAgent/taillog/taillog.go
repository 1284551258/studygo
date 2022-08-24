package taillog

import (
	"context"
	"fmt"
	"github.com/1284551258/logAgent/kafka"
	"github.com/hpcloud/tail"
)

type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewTailTask TailTask的构造函数
func NewTailTask(path string, topic string) *TailTask {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj := &TailTask{
		path:   path,
		topic:  topic,
		ctx:    ctx,
		cancel: cancel,
	}
	tailObj.init()
	return tailObj

}

func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail faile failed,err:", err)
	}
	go t.Run() //开启协程去读取日志发送到kafka
}
func (t *TailTask) Run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("删除任务%s_%s", t.topic, t.path)
			return
		case line := <-t.instance.Lines:
			//kafka.SendToKafka(t.topic, line.Text)//把日志发往kafka
			//把日志发往一个通道
			kafka.SendToChan(t.topic, line.Text)
			//kafka那个包中有一个单独的goroutine去取日志发到kafka
		}
	}
}

// ReadChan 从chan读取日志
func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}
