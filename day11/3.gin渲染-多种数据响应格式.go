package main

/*
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()
	//1.json格式
	r.GET("/someJson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "someJson", "status": 200})
	})
	//2.结构体格式
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "zhangsan"
		msg.Message = "someStruct"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})
	//3.xml格式
	r.GET("/someXml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "someXml", "status": 200})
	})

	//4.yaml格式
	r.GET("/someYaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "someYaml", "status": 200})
	})
	//5.protobuf格式
	r.GET("/someProtobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8000")
}

*/
