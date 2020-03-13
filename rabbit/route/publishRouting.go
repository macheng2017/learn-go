package main

import (
	"fmt"
	"learngo/rabbit/RabbitMQ"
	"strconv"
	"time"
)

// 路由模式最大的不同可以在生产端指定消费端的路由
func main() {
	macOne := RabbitMQ.NewRabbitMQRouting("exMac", "mac_one")
	macTwo := RabbitMQ.NewRabbitMQRouting("exMac", "mac_two")

	for i := 0; i <= 10; i++ {
		macOne.PublishRouting("hello mac one !" + strconv.Itoa(i))
		macTwo.PublishRouting("hello mac Two !" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
