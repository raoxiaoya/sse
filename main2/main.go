package main

import (
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
)

func main() {
	engin := gin.Default()

	engin.Any("/stream", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		// c.SSEvent("start", "start...")
		sse.Event{
			Id:    "1",
			Event: "start",
			Data:  "start...",
		}.Render(c.Writer)

		c.Writer.Flush()

		for i := 0; i < 10; i++ {
			sse.Event{
				Id:   "1",
				Data: "SSE data",
			}.Render(c.Writer)

			c.Writer.Flush() // 需要手动刷新输出缓冲区

			time.Sleep(1 * time.Second)
		}

		// c.SSEvent("end", "end...")
		sse.Event{
			Id:    "1",
			Event: "end",
			Data:  "end...",
		}.Render(c.Writer)
	})

	engin.Run(":8080")
}
