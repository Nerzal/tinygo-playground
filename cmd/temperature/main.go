package main

import (
	"machine"
)

func main() {
	println("starting temperature program")

	i2c := machine.I2C{}
	i2c.Configure(machine.I2CConfig{})

	waterLevel := machine.Pin(5)
	waterLevel.Configure(machine.PinConfig{machine.PinInput})

	for {
		println("trying to read data")

		// buffer := []byte{}
		// write := []byte{}

		// i2c.Tx(0, write, buffer)

		// println(string(buffer))
		// time.Sleep(time.Second)

		// tmpData := make([]byte, 2)

		// err := i2c.Bus.ReadRegister(0, 0, tmpData)
		// if err != nil {
		// 	println(err)
		// }

		// temperatureSum := int32((int16(tmpData[0])<<8 | int16(tmpData[1])) >> 4)
		// if (temperatureSum & int32(1<<11)) == int32(1<<11) {
		// 	temperatureSum |= int32(0xf800)
		// }

		// println(temperatureSum)
		// println(tmpData[0])
		// println("finished reading")
		// time.Sleep(time.Second)
	}

}
