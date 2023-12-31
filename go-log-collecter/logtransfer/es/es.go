package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

var (
	esClient  *elastic.Client
	logESChan chan LogData
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

// 初始化ES，准备接受KAFKA那边发出来的数据
func Init(address string, chan_max_size int, workers int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	esClient, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		return
	}
	fmt.Println("connect to es success")
	logESChan = make(chan LogData, chan_max_size)
	for i := 0; i < workers; i++ {
		go SendToES()
	}
	return
}

func SendToChan(data LogData) {
	logESChan <- data
}

// 发送数据到ES
func SendToES() {
	for {
		select {
		case msg := <-logESChan:
			// 链式操作
			put1, err := esClient.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
			if err != nil {
				// Handle error
				fmt.Printf("send to es failed, err: %v\n", err)
				continue
			}
			fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}
