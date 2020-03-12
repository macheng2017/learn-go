package main

import (
	"fmt"
	"learngo/rabbit/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitMQSimple := RabbitMQ.NewRabbitMQSimple("machengSimple")

	for i := 0; i <= 100; i++ {
		rabbitMQSimple.PublishSimple("hello world!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
