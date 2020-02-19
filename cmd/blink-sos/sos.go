package main

import (
	"time"

	"machine"
)

func main() {
	led := machine.Pin(12)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// while(true)
	for {
		blink3(led, 500, 350)

		blink3(led, 1500, 350)

		blink3(led, 500, 350)
	}
}

func blink3(led machine.Pin, duration, lowDuration time.Duration) {
	for i := 0; i < 3; i++ {
		led.High()
		time.Sleep(time.Millisecond * duration)
		led.Low()
		time.Sleep(time.Millisecond * lowDuration)
	}
}
