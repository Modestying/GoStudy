package main

import (
	"fmt"
	"time"
	"unsafe"

	log "github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
)

type PSubscribeCallback func(pattern, channel, message string)

type PSubscriber struct {
	client redis.PubSubConn
	cbMap  map[string]PSubscribeCallback
}

func (c *PSubscriber) PConnect(ip string, port uint16) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Critical("redis dial failed.")
	}

	c.client = redis.PubSubConn{conn}
	c.cbMap = make(map[string]PSubscribeCallback)

	go func() {
		for {
			log.Debug("wait...")
			switch res := c.client.Receive().(type) {
			case redis.Message:
				pattern := (*string)(unsafe.Pointer(&res.Pattern))
				channel := (*string)(unsafe.Pointer(&res.Channel))
				message := (*string)(unsafe.Pointer(&res.Data))
				c.cbMap[*channel](*pattern, *channel, *message)
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", res.Channel, res.Kind, res.Count)
			case error:
				log.Error("error handle...")
				continue
			}
		}
	}()

}
func (c *PSubscriber) Psubscribe(channel interface{}, cb PSubscribeCallback) {
	err := c.client.PSubscribe(channel)
	if err != nil {
		log.Critical("redis Subscribe error.")
	}

	c.cbMap[channel.(string)] = cb
}

func timeoutCallback(patter, chann, device string) {
	log.Debug("timeoutCallback patter : "+patter+" channel : ", chann, " offline device : ", device)
}

func main() {
	var psub PSubscriber
	psub.PConnect("127.0.0.1", 6397)
	psub.Psubscribe("__keyevent@0__:expired", timeoutCallback)
	for {
		time.Sleep(1 * time.Second)
	}
}
