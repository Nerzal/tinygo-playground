package main

import (
	"strconv"
	"syscall/js"
)


// Service is used as service
type Service interface {
}

type service struct {

}

// NewService instantiates a new Service
func NewService() Service {
   return &service{}
}












func main() {

}

//go:export add
func add(a, b int) int {
	return a + b
}

//go:export update
func update() {
	document := js.Global().Get("document")
	aStr := document.Call("getElementById", "a").Get("value").String()
	bStr := document.Call("getElementById", "b").Get("value").String()
	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)
	result := add(a, b)
	document.Call("getElementById", "result").Set("value", result)
}
