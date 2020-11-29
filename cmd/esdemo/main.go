package main

import (
	"context"
	"fmt"
	"os"

	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Printf("connect to es error:%v\n",err)
		os.Exit(1)
	}

	fmt.Println("connect to es success")
	p1 := Student{Name: "rion", Age: 22, Married: false}
	put1, err := client.Index().
		Index("student").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		fmt.Printf("save data to es error:%v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}