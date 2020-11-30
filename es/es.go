package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

type LogData struct {
	Topic string `json:"topic"`
	Obj   Msg    `json:"obj"`
}

type Msg struct {
	Data string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

func SendToES(msg *LogData) {
	ch <- msg
}

func Init(address string, chanSize, gSize int) error {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	var err error
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		fmt.Printf("connect to es error:%v\n", err)
		return err
	}
	ch = make(chan *LogData, chanSize)
	for i := 1; i <= gSize; i++ {
		go run()
	}
	return nil
}

func run() {
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().
				Index(msg.Topic).
				BodyJson(msg.Obj).
				Do(context.Background())
			if err != nil {
				fmt.Printf("save data to es error:%v\n", err)
				continue
			}
			fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}

}
