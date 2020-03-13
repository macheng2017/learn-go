package main

import "learngo/rabbit/RabbitMQ"

func main() {
	macTwo := RabbitMQ.NewRabbitMQRouting("exMac", "mac_two")
	macTwo.ReceiveRouting()
}
