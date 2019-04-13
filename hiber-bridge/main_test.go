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

func testInitConfiguration(t *testing.T) {
	// var config configfile
	dir = "./config.json"
	if initConfiguration(dir) == false {
		t.Errorf("Coiuld not parse config.json")
	}
}

func testFindSerialPort(t *testing.T) {
	if findSerialPort() == false {
		t.Errorf("Coiuld not find any serial port")
	}
}
