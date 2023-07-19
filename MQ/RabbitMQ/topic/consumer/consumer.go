package main

import (
	"log"
	"time"

	hkvision "github.com/Modestying/GoStudy/MQ/RabbitMQ/topic/hkvision/v1"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

const (
	// 连c接地址从控制台获取
	host = "amqp-xx.rabbitmq.xx.tencenttdmq.com"
	// 角色名称, 位于【角色管理】页面
	username = "admin"
	// 角色密钥, 位于【角色管理】页面
	password = "eyJrZXl..."
	// 要使用vhost全称
	vhost = "amqp-...|Vhost"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 创建连接
	conn, err := amqp.Dial("amqp://admin:admin508@192.168.10.210:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// 建立通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)

	// 声明消息队列
	q, err := ch.QueueDeclare(
		"topic_queue2", // name
		false,          // durable
		false,          // delete when unused
		true,           // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 绑定消息队列到指定交换机并订阅对应的 routing key
	var routing_key = [3]string{"common.*", "smart.*"}
	for i := range routing_key {
		err = ch.QueueBind(
			q.Name,          // queue name
			routing_key[i],  // routing key
			"alarm_message", // exchange
			false,
			nil,
		)
		failOnError(err, "Failed to bind a queue")
	}

	// 设置每次投递一个消息
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	// 创建消费者并消费指定消息队列中的消息
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // 设置为非自动确认(可根据需求自己选择)
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// 持续获取消息队列中的消息
	forever := make(chan bool)

	msg := &hkvision.AlarmInfo{}
	go func() {
		for d := range msgs {
			//log.Printf("Received a message: %s\n", d.Body)
			// t := time.Duration(1)
			// time.Sleep(t * time.Second)
			// // 手动回复ack
			// d.Ack(false)
			time.Sleep(time.Second * 1)
			if err := proto.Unmarshal(d.Body, msg); err != nil {
				log.Fatalln("Failed to parse message:", err)
			}
			log.Printf("Received a message: %s\n", msg)
		}
	}()

	log.Printf(" [Consumer1(topics: *.bbb.*/aaa.#/*.*.aaa)] Waiting for messages.")
	<-forever
}
