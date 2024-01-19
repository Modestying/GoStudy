package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp091.Dial("amqp://sipServer:123456@localhost:5672/shuGuo")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"asd",
		"direct",
		true,
		true,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	queue, err := ch.QueueDeclare(
		"asd", //队列名
		false, //是否持久化
		false, // 是否自动删除
		true,  // 是否独占
		false, // 是否阻塞
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		queue.Name,
		"wqe",
		"asd",
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")
	select {}
	return
	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
		msg.Ack(true)
	}
}
