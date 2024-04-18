package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	A string `json:"A" yaml:"A"`
	B string `json:"B" yaml:"B"`
	C int    `json:"C,string"`
}

func Json() {
	config := Config{}
	x, _ := os.Open("config.json")
	defer x.Close()
	decoder := json.NewDecoder(x)
	decoder.Decode(&config)
	fmt.Println(config)

}

func Yaml() {
	conf := Config{}
	if f, err := os.Open("config.yaml"); err != nil {
	} else {
		yaml.NewDecoder(f).Decode(&conf)
	}
	fmt.Println(conf)
}

func JsonNumber() {
	cfg := map[string]int{
		"C": 12,
	}
	bytes, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	var c Config
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", c)
	_, err = json.Marshal(c)
	if err != nil {
		panic(err)
	}
}

func main() {
	JsonNumber()
}
