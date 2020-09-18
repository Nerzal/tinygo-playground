package main

import (
	"machine"
	"time"

	"github.com/tinygo-org/drivers/hd44780i2c"
)

// change these to test a different UART or pins if available
var (
	uart = machine.UART0
	tx   = machine.UART_TX_PIN
	rx   = machine.UART_RX_PIN
)

//   The circuit:
//  * LCD RS pin to digital pin 4
//  * LCD Enable pin to digital pin 11
//  * LCD D4 pin to digital pin 13
//  * LCD D5 pin to digital pin 12
//  * LCD D6 pin to digital pin 11
//  * LCD D7 pin to digital pin 10
//  * LCD VSS pin to ground
//  * LCD VCC pin to 5V
//  * LCD K pin to 5V
//  * LCD A pin to 3.3V
//  * 10K resistor:
//  * ends to +5V and ground
//  * wiper to LCD VO pin (pin 3)
func main() {
	println("starting program")

	uart.Configure(machine.UARTConfig{TX: tx, RX: rx})
	uart.Write([]byte("Echo console enabled. Type something then press enter:\r\n"))

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	lcd := hd44780i2c.New(machine.I2C0, 0x27) // some modules have address 0x3F

	lcd.Configure(hd44780i2c.Config{
		Width:       16, // required
		Height:      2,  // required
		CursorOn:    true,
		CursorBlink: true,
	})

	lcd.Print([]byte(" TinyGo\n  LCD Test "))

	// CGRAM address 0x0-0x7 can be used to store 8 custom characters
	lcd.CreateCharacter(0x0, []byte{0x00, 0x11, 0x0E, 0x1F, 0x15, 0x1F, 0x1F, 0x1F})
	lcd.Print([]byte{0x0})

	input := make([]byte, 64)
	i := 0
	for {
		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()

			switch data {
			case 13:
				// return key
				uart.Write([]byte("\r\n"))
				uart.Write([]byte("You typed: "))
				uart.Write(input[:i])
				uart.Write([]byte("\r\n"))

				if string(input[:i]) == "reset" {
					lcd.ClearDisplay()
					lcd.Print([]byte(" TinyGo\n  LCD Test "))

					lcd.CreateCharacter(0x0, []byte{0x00, 0x11, 0x0E, 0x1F, 0x15, 0x1F, 0x1F, 0x1F})
					lcd.Print([]byte{0x0})
					continue
				}

				lcd.ClearDisplay()
				lcd.Print(input[:i])
				i = 0
			default:
				// just echo the character
				uart.WriteByte(data)
				input[i] = data
				i++
			}
		}
		time.Sleep(10 * time.Millisecond)
	}

	// uart.Configure(machine.UARTConfig{TX: tx, RX: rx})
	// println("uart configured")

	// for {
	// 	data := make([]byte, 8)

	// 	i, err := uart.Read(data)
	// 	if err != nil {
	// 		println(err)
	// 	}

	// 	if i == 0 {
	// 		continue
	// 	}

	// 	lcd.ClearDisplay()
	// 	lcd.SetCursor(2, 1)

	// 	lcd.Print(data)
	// }

	// You can use https://maxpromer.github.io/LCD-Character-Creator/
	// to crete your own characters.

	// time.Sleep(time.Millisecond * 7000)

	// for i := 0; i < 5; i++ {
	// 	lcd.BacklightOn(false)
	// 	time.Sleep(time.Millisecond * 250)
	// 	lcd.BacklightOn(true)
	// 	time.Sleep(time.Millisecond * 250)
	// }

	// lcd.CursorOn(false)
	// lcd.CursorBlink(false)

	// i := 0
	// for {

	// 	lcd.ClearDisplay()
	// 	lcd.SetCursor(2, 1)
	// 	lcd.Print([]byte(strconv.FormatInt(int64(i), 10)))
	// 	i++
	// 	time.Sleep(time.Millisecond * 100)

	// }
}
