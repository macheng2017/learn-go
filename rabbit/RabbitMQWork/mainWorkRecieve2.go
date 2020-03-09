package main

import "learngo/rabbit/RabbitMQ"

func main() {
	rabbitMQSimple := RabbitMQ.NewRabbitMQSimple(
		"machengSimple")

	rabbitMQSimple.ConsumeSimple()

}
