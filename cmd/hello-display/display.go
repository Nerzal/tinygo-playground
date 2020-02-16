package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/hd44780"
)

//   The circuit:
//  * LCD RS pin to digital pin 12
//  * LCD Enable pin to digital pin 11
//  * LCD D4 pin to digital pin 5
//  * LCD D5 pin to digital pin 4
//  * LCD D6 pin to digital pin 3
//  * LCD D7 pin to digital pin 2
//  * LCD VSS pin to ground
//  * LCD VCC pin to 5V
//  * 10K resistor:
//  * ends to +5V and ground
//  * wiper to LCD VO pin (pin 3)
func main() {
	led := machine.Pin(13)
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	lcd, err := hd44780.NewGPIO4Bit(
		[]machine.Pin{machine.Pin(2), machine.Pin(3), machine.Pin(4), machine.Pin(5)}, //DataPins
		machine.Pin(11), // e pin
		machine.Pin(12), // RS Pin
		machine.Pin(8),  // RW Pin
	)

	if err != nil {
		led.High()
		for {
		}
	}

	lcd.Configure(hd44780.Config{
		Width:       16,
		Height:      2,
		CursorOnOff: true,
		CursorBlink: true,
	})

	for {
		led.High()
		time.Sleep(time.Millisecond * 250)

		lcd.Write([]byte("TinyGo"))
		lcd.Display()

		led.Low()
		time.Sleep(time.Millisecond * 250)

		lcd.Write([]byte("Hello World"))
		lcd.Display()
	}
}
