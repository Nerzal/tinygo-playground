package main

import (
	"time"

	"machine"
)

func main() {
	pingConfig := machine.PinConfig{Mode: machine.PinOutput}

	buttonIn := machine.Pin(2)
	buttonIn.Configure(machine.PinConfig{Mode: machine.PinInput})

	greenLED := machine.Pin(12)
	greenLED.Configure(pingConfig)

	yellowLED := machine.Pin(11)
	yellowLED.Configure(pingConfig)

	redLED := machine.Pin(10)
	redLED.Configure(pingConfig)

	testLEDs(redLED, yellowLED, greenLED)

	sleep(2000)
	// Traffic-Lights
	// RED
	// RED-YELLOW
	// YELLOW
	// GREEN
	// YELLOW
	// RED
	foo := make(chan bool, 1)

	go checkButton(buttonIn, foo)

	for {
		if <-foo {
			redLED.Low()
			yellowLED.Low()
			redLED.Low()
			sleep(10000)
		}

		redLED.High()
		sleep(1000)
		yellowLED.High()
		sleep(1000)
		redLED.Low()
		sleep(1000)
		yellowLED.Low()
		greenLED.High()
		sleep(1000)
		greenLED.Low()
		yellowLED.High()
		sleep(1000)
		yellowLED.Low()
	}
}

func checkButton(button machine.Pin, foo chan bool) {
	for {
		if button.Get() {
			foo <- true
		} else {
			foo <- false
		}
	}
}

func sleep(duration time.Duration) {
	time.Sleep(time.Millisecond * duration)
}

func testLEDs(redLED, yellowLED, greenLED machine.Pin) {
	greenLED.High()
	time.Sleep(time.Millisecond * 500)
	greenLED.Low()

	yellowLED.High()
	time.Sleep(time.Millisecond * 500)
	yellowLED.Low()

	redLED.High()
	time.Sleep(time.Millisecond * 500)
	redLED.Low()
}
