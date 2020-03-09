package main

import (
	"fmt"
	"learngo/rabbit/RabbitMQ"
)

func main() {
	rabbitMQSimple := RabbitMQ.NewRabbitMQSimple("machengSimple")

	rabbitMQSimple.PublishSimple("hello world!")
	fmt.Printf("发送成功")

}
