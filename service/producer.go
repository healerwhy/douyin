package main

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"time"
)

type KqConfig struct {
	Brokers []string
	Topic   string
}

type UserFavoriteOptMessage struct {
	OptStatus int64  `json:"payStatus"`
	Opt       string `json:"orderSn"`
}

func main() {
	var kqConfig KqConfig
	kqConfig.Brokers = []string{"127.0.0.1:9092"}
	kqConfig.Topic = "UserFavoriteOptService-topic"

	var pusher *kq.Pusher
	pusher = kq.NewPusher(kqConfig.Brokers, kqConfig.Topic)
	var msg UserFavoriteOptMessage
	for i := 0; i < 10; i++ {
		msg.OptStatus = 1
		msg.Opt = "hello"
		tmp, _ := json.Marshal(msg)
		e := pusher.Push(string(tmp))
		fmt.Println("push msg:", string(tmp))
		if e != nil {
			fmt.Println("push err:", e)
		}
		time.Sleep(time.Second * 2)
		msg.OptStatus = 2
		msg.Opt = "world"
		tmp, _ = json.Marshal(msg)
		e = pusher.Push(string(tmp))
		fmt.Println("push msg:", string(tmp))
		if e != nil {
			fmt.Println("push err:", e)
		}
		time.Sleep(time.Second * 2)
	}
}
