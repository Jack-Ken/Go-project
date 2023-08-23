package go_log_collecter

import (
	"fmt"
	"go-log-collecter/conf"
	producer "go-log-collecter/logagent/kafka"
	"go-log-collecter/logtransfer/es"
	consumer "go-log-collecter/logtransfer/kafka"
	"log"
	"strings"
)

func main() {
	// 1. 初始化kafka producer连接
	err := producer.InitProducer(strings.Split(conf.C.Kafka.Address, ";"), conf.C.Kafka.ChanMaxSize)
	if err != nil {
		log.Println("init kafka producer failed, err:%v\n", err)
		return
	}
	log.Println("init kafka success.")

	// 2. 初始化kafka consumer连接
	err = consumer.InitConsumer(strings.Split(conf.C.Kafka.Address, ";"))
	if err != nil {
		log.Println("init kafka consumer failed, err:%v\n", err)
		return
	}

	// 3. 从kafka取日志数据并放入channel中
	consumer.Run()

	// 1. 初始化ES
	// 1.1 初始化一个ES连接的client
	err = es.Init(conf.C.Es.Address, conf.C.Es.ChanMaxSize, conf.C.Es.Workers)
	if err != nil {
		fmt.Printf("init ES client failed,err:%v\n", err)
		return
	}
	fmt.Println("init ES client success.")
}
