package main

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	//1.建立连接和通信信道
	conn, err := amqp.Dial("amqp://user:12345@0.0.0.0:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)

	// 建立通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//2.声明交换机，类型和名称需要和存在的交换机保持一致
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "declare queue")
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "register consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var stu2 Student
			err := json.Unmarshal(d.Body, &stu2)
			if err != nil {
				fmt.Printf("反序列化错误 err=%v\n", err)
				return
			}
			fmt.Printf("反序列化后: Student=%v, Name=%v\n", stu2, stu2.Name)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

type Student struct {
	Name  string
	Age   int
	Skill string // 也可以不指定 tag标签，默认就是 变量名称
}
