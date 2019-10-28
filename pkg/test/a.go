package main

import "fmt"

var global int = 1

func hex(arg int) string {
	return fmt.Sprintf(`%b`, arg)
}

func main() {
	fmt.Println(hex(global))
}
