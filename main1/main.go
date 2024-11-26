package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/antage/eventsource.v1"
)

func main() {
	es := eventsource.New(nil, nil)
	defer es.Close()

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.Handle("/events", es)
	go func() {
		for {
			// 每2秒发送一条当前时间消息，并打印对应客户端数量
			// 自定义event,例如拼接上用户id，可以给用户单独推送
			es.SendEventMessage(fmt.Sprintf("hello, 现在的时间是: %s", time.Now()), "abc", "")
			
			// 默认的event是message是可以不写的
			// es.SendEventMessage(fmt.Sprintf("hello, 现在的时间是: %s", time.Now()), "", "")
			log.Printf("在线人数: %d", es.ConsumersCount())
			time.Sleep(2 * time.Second)
		}
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
