package main_test

import (
	"testing"

	zinnet "github.com/zinsoldat/zinnet-go"
)

func TestCreateServer(t *testing.T) {
	serverConfig := &zinnet.ServerConfig{
		Port: 8888,
		Host: "127.0.0.1",
	}
	zinnet.NewServer(serverConfig)

	// server should be creatable.
}
