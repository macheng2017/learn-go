package main

import "learngo/rabbit/RabbitMQ"

func main() {
	rabbitMQ := RabbitMQ.NewRabbitMQPubSub("newProduct")
	rabbitMQ.ReceiveSub()

}
