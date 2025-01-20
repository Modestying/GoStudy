package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func main() {
	data, _ := json.Marshal(map[string]any{
		"passwd": "123",
		"name":   "xx",
	})
	transport := &http.Transport{}

	for i := 0; i < 1000; i++ {
		cli := http.Client{
			Transport: transport,
		}
		req, err := cli.Post("http://localhost:8085/login", "application/json", bytes.NewReader(data))
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()
		io.Copy(io.Discard, req.Body)
	}

}
