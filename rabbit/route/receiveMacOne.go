package main

import "learngo/rabbit/RabbitMQ"

func main() {
	macOne := RabbitMQ.NewRabbitMQRouting("exMac", "mac_one")
	macOne.ReceiveRouting()
}
