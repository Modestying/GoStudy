package main

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "wss://192.168.1.196:59300", &websocket.DialOptions{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	defer c.CloseNow()
	time.Sleep(time.Second * 5)
}
