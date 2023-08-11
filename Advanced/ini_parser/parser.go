package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Config struct {
	Name    string `ini:"name"`
	Version int    `ini:"version"`
	Author  string `ini:"author"`
}

func parseIniFile(data string, config interface{}) error {
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			section := line[1 : len(line)-1]
			v := reflect.ValueOf(config).Elem()
			field := v.FieldByName(section)
			if !field.IsValid() {
				return fmt.Errorf("Invalid section: %s", section)
			}
			if field.Kind() != reflect.Struct {
				return fmt.Errorf("Invalid section type: %v", field.Kind())
			}
			parseSection(lines, field.Addr().Interface())
		} else {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				return fmt.Errorf("Invalid line: %s", line)
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if len(key) == 0 || len(value) == 0 {
				return fmt.Errorf("Invalid line: %s", line)
			}
			field := reflect.ValueOf(config).Elem().FieldByName(key)
			if !field.IsValid() {
				return fmt.Errorf("Invalid key: %s", key)
			}
			field.SetString(value)
		}
	}
	return nil
}

func parseSection(lines []string, config interface{}) {
	fieldPtr := reflect.ValueOf(config).Elem()
	fieldType := fieldPtr.Type()
	for i := 0; i < fieldType.NumField(); i++ {
		field := fieldType.Field(i)
		if tag, ok := field.Tag.Lookup("ini"); ok {
			value := findValue(lines, tag)
			if len(value) > 0 {
				fieldPtr.Field(i).SetString(value)
			}
		}
	}
}

func findValue(lines []string, key string) string {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		configKey := strings.TrimSpace(parts[0])
		configValue := strings.TrimSpace(parts[1])
		if configKey == key {
			return configValue
		}
	}
	return ""
}

func main() {
	data := `
name = MyApp

[info]
version = 1
author = John Doe

[database]
host = localhost
port = 3306
`
	config := &Config{}
	err := parseIniFile(data, config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Name:", config.Name)
	fmt.Println("Version:", config.Version)
	fmt.Println("Author:", config.Author)
	fmt.Println("Host:", config.info.Host)
	fmt.Println("Port:", config.info.Port)
}
