package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

//ALTER USER ‘root’@‘localhost’ IDENTIFIED WITH caching_sha2_password BY '0Gu7KxSFYBtBWzDI';

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// /etc/my.cnf /etc/mysql/my.cnf /opt/homebrew/etc/my.cnf ~/.my.cnf

func main() {
	conn, err := amqp091.Dial("amqp://sipServer:123456@localhost:5672/shuGuo")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"asd",    // 交换机名，应该以"amq."开头
		"direct", // 交换机类型, fanout(信息发布给所有队列) direct(信息只发布给路由完全匹配的队列), topic(*.特殊路由), headers(根据消息头来路由)
		true,     // 是否持久化，影响重启后是否存在
		true,     // 是否自动删除，当最后一个绑定到交换机上的队列删除后，自动删除交换机
		false,    // 是否内置，内置是rabbitmq内部使用的，默认false
		false,    // 是否等待服务器响应，true等待，false不等待,就是直接用，不等待mq回复
		nil)
	if err == nil {
		fmt.Println("declare success")
	} else {
		fmt.Println("err:", err.Error())
	}
	select {}
	return
	failOnError(err, "Failed to declare an exchange")

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	body := "Hello World!"
	err = ch.PublishWithContext(
		ctx,
		"logs",
		"",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
