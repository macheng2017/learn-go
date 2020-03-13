package main

import "learngo/rabbit/RabbitMQ"

func main() {
	macOne := RabbitMQ.NewRabbitMTopic("exMacTopic", "#")
	macOne.ReceiveTopic()
}
