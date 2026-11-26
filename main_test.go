package main

import "testing"

func TestRuntime(t *testing.T) {
	if 1+1 != 2 { t.Fatal("runtime check failed") }
}
