package RabbitMQ

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// [streadway/amqp: Go client for AMQP 0.9.1](https://github.com/streadway/amqp)
//const MQURL = "amqp://guest:guest@192.168.99.100:5672/macheng"
const MQURL = "amqp://guest:guest@127.0.0.1:5672/macheng"

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

// step1
// 根据基础实例,通过传入不同的参数创建不同模式的实例
// 简单模式下rabbitmq实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	return rabbitmq
}

// step2 简单模式下生产者

func (r *RabbitMQ) PublishSimple(message string) {
	// 申请队列,如果队列不存在则创建,若存在则直接用
	// 保证队列存在,消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性(仅自己可见)
		false,
		// 是否阻塞(是否等待响应)
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	// 2 发送消息到队列中
	r.channel.Publish(r.Exchange, r.QueueName,
		// 若true 根据exchange类型,和routkey规则如果无法找到符合条件的队列,将会退回消息给发送者
		false,
		//若为true,当exchange发送消息到队列发现队列没有绑定消费者,则返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (r *RabbitMQ) ConsumeSimple() {
	// 申请队列,如果队列不存在则创建,若存在则直接用
	// 保证队列存在,消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性(仅自己可见)
		false,
		// 是否阻塞(是否等待响应)
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	// 接收消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		// 用来区分多个消费者
		"",
		// 是否自动应答
		true,
		//是否具有排他性
		false,
		// 如果设置为true,表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		// 是否阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)

	// 开启协程处理消息
	go func() {
		for d := range msgs {
			// 实现我们要处理的逻辑函数
			log.Printf("Received a message: %s", d.Body)
			//fmt.Printf("%v", d)

		}
	}()

	log.Printf("[*] Waiting for message, To exit press CTRL + C ")
	<-forever

}

//订阅模式下创建RabbitMQ实例
func NewRabbitMQPubSub(exchangName string) *RabbitMQ {
	rabbitMQ := NewRabbitMQ("", exchangName, "")
	var err error
	// 获取连接
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.Mqurl)
	rabbitMQ.failOnErr(err, "failed to connect rabbitmq")
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.failOnErr(err, "failed to open a channel")
	return rabbitMQ
}

// 订阅模式下生产
//[RabbitMQ tutorial - Publish/Subscribe — RabbitMQ](https://www.rabbitmq.com/tutorials/tutorial-three-go.html)
func (r *RabbitMQ) PublishPub(message string) {
	// 1.声明一个交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		// 定义交换机的类型 广播类型
		"fanout",
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息,仅用来进行exchanges之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "failed to declare an exchange")
	// 2. 发送消息
	r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 订阅模式下消费

func (r *RabbitMQ) ReceiveSub() {
	// 1.声明一个交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		// 定义交换机的类型 广播类型
		"fanout",
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息,仅用来进行exchanges之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "failed to declare an exchange")
	//2. 声明一个队列
	q, err := r.channel.QueueDeclare(
		// "" 随机生成队列
		"",
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性(仅自己可见)
		false,
		// 是否阻塞(是否等待响应)
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		log.Println(err)
	}
}
