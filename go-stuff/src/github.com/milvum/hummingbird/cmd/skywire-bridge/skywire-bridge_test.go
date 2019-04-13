package main

import (
	"testing"
)

// example test
func TestGetPath(t *testing.T) {
	dir := getPath()
	if dir == "" {
		t.Errorf("Path was expected, got nothing, want non empty dir %s", dir)
	}
}
