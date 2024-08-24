package main

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Kafka struct {
		Brokers []string "mapstructure:'brokers'"
		Topic   string   "mapstructure:'topic'"
	}

	Log struct {
		Filepath string "mapstructure:'filepath'"
	}
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	//var config Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	//if err := viper.Unmarshal(&config); err != nil {
	//	log.Fatalf("can not parse config file, %s", err)
	//}
	//
	//fmt.Printf("%#v", config)

	// kafka配置
	brokers := viper.GetStringSlice("kafka.brokers")
	topic := viper.GetString("kafka.topic")

	//conf :=
}
