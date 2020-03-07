package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// [streadway/amqp: Go client for AMQP 0.9.1](https://github.com/streadway/amqp)
const MQURL = "amqp://macheng:machenguser@192.168.99.100:5672/macheng"

type RabbitMQ struct {
	conn *amqp.Connection

	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	Key      string
	// 链接信息
	Mqurl string
}

// 结构体的函数 即 对象的函数
// 断开channel和connection
// 这种在函数前面加上()并声明一个结构体类型的写法可以理解为this,
// 这个是go语言定义对象中(结构体)定义方法的方式
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 创建结构体基础实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	//下面就是上面定义的结构体,有点类似json对象
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
	var err error
	// /创建链接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建链接错误")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq
}

// 根据基础实例,通过传入不同的参数创建不同模式的实例
// 简单模式下rabbitmq实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	return rabbitmq
}
