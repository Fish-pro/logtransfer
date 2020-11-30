package kafka

import (
	"fmt"
	"github.com/Fish-pro/logtransfer/es"
	"github.com/Shopify/sarama"
)

type LogData struct {
	Data string `json:"data"`
}

func Init(address []string, topic string) error {
	consumer, err := sarama.NewConsumer(address, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return err
	}
	fmt.Println("分区列表：", partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			defer pc.AsyncClose()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%s\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				// 直接给es
				ld := LogData{
					Data: string(msg.Value),
				}
				err := es.SendToES(topic, ld)
				if err != nil {
					fmt.Printf("send to es error:%v\n", err)
					continue
				}
			}
		}(pc)
	}
	return nil
}
