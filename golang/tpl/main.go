package main

import (
	_ "embed"
	"os"
	"text/template"
)

//go:embed service.tpl
var TplFile string

type Service struct {
	Name        string
	Path        string
	Config      string
	Description string
}

func main() {
	tmpl, err := template.New("tpl").Parse(TplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, Service{
		Name:        "test",
		Path:        "/test",
		Config:      "test.yaml",
		Description: "test service",
	})

	if err != nil {
		panic(err)
	}

}
