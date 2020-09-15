package main

import "time"

func main() {
	for {
		println("Hello World")
		time.Sleep(500 * time.Millisecond)
	}
}
