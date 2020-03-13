package main

import "learngo/rabbit/RabbitMQ"

func main() {
	macTwo := RabbitMQ.NewRabbitMTopic("exMacTopic", "mac.*.two")
	macTwo.ReceiveTopic()
}
