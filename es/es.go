package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
)

var client *elastic.Client

func Init(address string) error {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	var err error
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		fmt.Printf("connect to es error:%v\n", err)
		return err
	}
	return nil
}

func SendToES(indexStr string, data interface{}) error {
	put1, err := client.Index().
		Index(indexStr).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		fmt.Printf("save data to es error:%v\n", err)
		return err
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return nil
}
