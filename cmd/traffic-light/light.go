package main

import (
	"time"

	"machine"
)

func main() {
	pingConfig := machine.PinConfig{Mode: machine.PinOutput}

	greenLED := machine.Pin(12)
	greenLED.Configure(pingConfig)

	yellowLED := machine.Pin(11)
	yellowLED.Configure(pingConfig)

	redLED := machine.Pin(10)
	redLED.Configure(pingConfig)

	testLEDs(redLED, yellowLED, greenLED)

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
