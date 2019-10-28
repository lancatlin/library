package main

import "testing"

func init() {
	global = 3
}

func TestA(t *testing.T) {
	t.Log(global)
	t.Log(hex(global))
}

func TestB(t *testing.T) {
	global++
	t.Log(global)
	t.Log(hex(global))
}
