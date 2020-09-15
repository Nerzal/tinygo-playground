package main

import (
	"machine"
	"runtime/volatile"
	"time"
)

const COMPLETE_IN_WATER uint16 = 20000
const DRY_THRESHOLD uint16 = 64
const MOIST_THRESHOLD uint16 = 10000

// Connect S Pin to ADC5
// Connect + to 5V
// Conntect - to GND
func main() {
	machine.InitADC()

	register := volatile.Register16{}
	register.Set(10)

	pump := machine.Pin(machine.PD7)
	pump.Configure(machine.PinConfig{Mode: machine.PinOutput})

	waterLevel := machine.ADC{machine.ADC5}
	waterLevel.Configure()

	soilSensor := machine.ADC{machine.ADC3}
	soilSensor.Configure()

	println("pump")
	pump.High()
	time.Sleep(5 * time.Second)
	pump.Low()

	for {
		val := soilSensor.Get()

		// > 30.000 dry / in air
		// < 19.000 moist
		// < 10.000 very moist
		println("soil sensor val: ", val)
		time.Sleep(1000 * time.Millisecond)
	}

}

func getWaterLevel(sensor machine.ADC) uint16 {
	val := sensor.Get()
	println("water level value : ", val)

	if val <= MOIST_THRESHOLD {
		println("Not in water, but moist")
	} else if val <= DRY_THRESHOLD {
		println("Dry, no water detected")
	} else if val > MOIST_THRESHOLD {
		println("Water detected")
	} else if val >= COMPLETE_IN_WATER {
		println("Completely in water")
	}

	return val
}
