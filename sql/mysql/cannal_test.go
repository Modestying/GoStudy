package main

import (
	"testing"

	"github.com/go-mysql-org/go-mysql/canal"
)

func TestNewMyEventHandler(t *testing.T) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "localhost:3306"
	cfg.User = "root"
	cfg.Password = "fyl183279"
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "demo"
	cfg.Dump.Tables = []string{"movie"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		t.Fatal(err)
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	c.Run()
}
