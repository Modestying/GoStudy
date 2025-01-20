package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

var presence = "http://localhost:9090/get"

func main() {

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		DisableKeepAlives: false,            // 启用 Keep-Alive
		MaxIdleConns:      100,              // 最大空闲连接数
		IdleConnTimeout:   90 * time.Second, // 空闲连接超时时间
	}

	for i := 0; i < 100; i++ {
		cli := http.Client{
			Transport: tr,
			Timeout:   time.Second,
		}
		req, err := http.NewRequest(http.MethodPost, presence, nil)
		if err != nil {
			panic(err)
		}
		resp, err := cli.Do(req)
		if err != nil {
			panic(err)
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(data))
		_ = resp.Body.Close()
	}

}
