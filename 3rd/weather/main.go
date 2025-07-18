package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const HfUrl = "http://59.204.85.129:58080/qweather/"

func main() {
	req := url.Values{
		"location": []string{"101230505"},
		"key":      []string{"3f0a6d095e794f33b661644ca888dd9d"},
	}
	resp, err := getHfData(HfUrl+"weather/now", req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	fmt.Println("请求成功:", string(resp))
}

func getHfData(action string, values url.Values) ([]byte, error) {
	hfUrl, urlErr := url.Parse(action + "?" + values.Encode())
	if urlErr != nil {
		return nil, urlErr
	}
	c := &http.Client{
		Timeout: 2 * time.Second,
	}
	resp, getErr := c.Get(hfUrl.String())
	if getErr != nil {
		return nil, getErr
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取和风天气请求失败！" + err.Error())
		return nil, errors.New("读取和风天气请求失败！" + err.Error())
	}
	return body, nil
}
