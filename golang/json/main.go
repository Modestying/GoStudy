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

func main() {
	Json()
	Yaml()
}
