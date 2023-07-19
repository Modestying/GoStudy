package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Demo struct {
	name string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func main() {
	conn, err := amqp.Dial("amqp://user:12345@0.0.0.0:5672/")
	failOnError(err, "build connection")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable,持久化,是否存盘
		false,   // delete when unused 是否自动删除，前提是有消费者连接到这个队列，之后所有消费者都断开连接
		false,   // exclusive 是否具有排他性，如果是true，那么只有创建这个队列的连接才能使用这个队列，连接关闭后队列自动删除
		false,   // no-wait 是否阻塞
		nil,     // arguments
	)
	failOnError(err, "declare a queue")
	for i := 0; i < 10; i++ {
		//body := "hello dd"
		var body string
		body = `{"Name":"tom","Age":12,"Skill":"football"}`
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(body),
			},
		)
		failOnError(err, "failed send msg")
	}

}
