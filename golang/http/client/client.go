package main

import (
	"context"
	"net"
	"net/http"
)

var Conn net.Conn

func main() {
	net.ParseIP("")
	Conn, _ = net.DialTCP("tcp",
		&net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 22222,
		},
		&net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8080,
		},
	)

	tr := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return Conn, nil
		},
	}

	for i := 0; i < 100; i++ {
		cli := http.Client{
			Transport: tr,
		}
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/get", nil)
		if err != nil {
			panic(err)
		}
		resp, err := cli.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	}
	Conn.Close()

}
