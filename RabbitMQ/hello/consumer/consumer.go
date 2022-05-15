package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var (
	username = "mimo"
	password = "mimo431"
	host     = "118.178.180.57"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	//1.建立连接和通信信道
	conn, err := amqp.Dial("amqp://" + username + ":" + password + "@" + host + ":5672/")
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

	//2.声明交换机，类型和名称需要和存在的交换机保持一致
	err = ch.ExchangeDeclare(
		"test_direct",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Declare exchange")

	//3.发送消息
	body := "test message: direct"
	err = ch.Publish(
		"test_direct",
		"queue3",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Send message to exchange")

	//err = ch.Publish("",
	//	"queue3",
	//	false,
	//	false,
	//	amqp.Publishing{
	//		ContentType: "text/plain",
	//		Body:        []byte("send to queue"),
	//	},
	//)
	//failOnError(err, "Send message to queue3")

}
