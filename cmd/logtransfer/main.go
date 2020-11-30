package main

import (
	"fmt"
	"github.com/Fish-pro/logtransfer/config"
	"github.com/Fish-pro/logtransfer/es"
	"github.com/Fish-pro/logtransfer/kafka"
	"gopkg.in/ini.v1"
	"os"
)

func main() {
	// 0.加载配置文件
	var conf config.LogTransfer
	err := ini.MapTo(&conf, "./config/config.ini")
	if err != nil {
		fmt.Printf("load config file failed error:%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("config file:%v\n", conf)

	// 1.初始化es
	err = es.Init(conf.Es.Address)
	if err != nil {
		fmt.Printf("init es failed,err:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("init es success")

	// 2.初始化kafka
	err = kafka.Init([]string{conf.Kafka.Address}, conf.Kafka.Topic)
	if err != nil {
		fmt.Printf("init kafka consumer failed,err:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("init kafka success")
	select {}
}
