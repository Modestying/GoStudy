package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	reader *kafka.Reader
	topic  = "demo"
)

func writeKafka(ctx context.Context) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(":9095"),
		Topic:                  topic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireAll,
		AllowAutoTopicCreation: true,
		MaxAttempts:            5,
	}
	defer writer.Close()
	for i := 0; i < 3; i++ {
		if err := writer.WriteMessages(
			ctx,
			kafka.Message{Key: []byte("1"), Value: []byte(time.Now().Format("2022-01-02 12:21:12") + " a的")},
			kafka.Message{Key: []byte("2"), Value: []byte(time.Now().Format("2022-01-02 12:21:12") + " q的")},
			kafka.Message{Key: []byte("3"), Value: []byte(time.Now().Format("2022-01-02 12:21:12") + " e的")},
			kafka.Message{Key: []byte("2"), Value: []byte(time.Now().Format("2022-01-02 12:21:12") + " w的")},
		); err != nil {
			if err == kafka.LeaderNotAvailable {
				time.Sleep(time.Millisecond * 500)
				continue
			} else {
				fmt.Println("信息写入失败:", err)
			}
		} else {
			break
		}
	}
}

func readKafka(ctx context.Context) {
	reader := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers:        []string{":9095"},
			Topic:          topic,
			CommitInterval: 1 * time.Second,
			GroupID:        "rec_team",
			StartOffset:    kafka.FirstOffset,
		},
	)
	for {
		if msg, err := reader.ReadMessage(ctx); err != nil {
			fmt.Println("read failed ", err)
		} else {
			fmt.Printf("topic:%s,partition:%d,offset:%d,key=%s ,value=%s\n ",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		}
	}
}

func Listen() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

}
func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
	return
	writeKafka(context.Background())
	readKafka(context.Background())
}
