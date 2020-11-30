package config

type LogTransfer struct {
	Kafka KafkaConfig `ini:"kafka"`
	Es    EsConfig    `ini:"es"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type EsConfig struct {
	Address string `ini:"address"`
}
