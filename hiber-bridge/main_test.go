package main

import (
	"testing"
)

func TestGetPath(t *testing.T) {
	dir := getPath()
	if dir == "" {
		t.Errorf("Path was expected, got nothing, want non empty dir %s", dir)
	}
}

func TestHiberCommand(t *testing.T) {
	go hiberCommand("teststring")
	msg := <-d
	if msg == "" {
		t.Errorf("Path was expected, got nothing, want non empty dir %s", msg)
	}
}

// func TestOpenSerialPort(t *testing.T) {
// 	var config configfile
// 	dir = "./config.json"
// 	initConfiguration(config, dir)
// 	port := openSerialPort()
// 	if reflect.TypeOf(port) != nil {
// 		t.Errorf("Path was expected, got nothing, want non empty dir %s", port)
// 	}
// }
