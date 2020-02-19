package main

import (
	"syscall/js"
	"time"
)

var initialized bool = false
var timeStart time.Time

func main() {
	// go fml()
	// go gunther()
	document := js.Global().Get("document")

	if !initialized {
		message := getElementById(document, "message")
		message.Set("value", "")
		initialized = true
		timeStart = time.Now()
		go updateTimer(document)
	}

	message := getElementById(document, "message").Get("value")

	println(message.String())

	messageList := getElementById(document, "messagesList")
	il := document.Call("createElement", "il")
	il.Set("innerHTML", message)
	classList := il.Get("classList")
	classList.Call("add", "message-list-message")
	messageList.Call("appendChild", il)

}

func updateTimer(document js.Value) {
	for {
		now := time.Now()
		diff := now.Sub(timeStart)
		out := time.Time{}.Add(diff)

		timer := getElementById(document, "presentationTimer")
		timer.Set("innerHTML", out.Format("15:04:05"))
		time.Sleep(500 * time.Millisecond)
	}
}

func getElementById(document js.Value, name string) js.Value {
	return document.Call("getElementById", name)
}

func fml() {
	for {
		println("FML!")
		time.Sleep(1250 * time.Millisecond)
	}
}

func gunther() {
	for {
		println("DAMN U GUNTHER! DON'T U DARE TAKE MY TABLE!!!!!")
		time.Sleep(750 * time.Millisecond)
	}
}
