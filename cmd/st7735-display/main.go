package main

func main() {

	// machine.SPI.Configure(machine.SPIConfig{
	// 	SCK:       machine.SPI1_SCK_PIN,
	// 	SDO:       machine.SPI1_SDO_PIN,
	// 	SDI:       machine.SPI1_SDI_PIN,
	// 	Frequency: 8000000,
	// })
	// machine.I2C0.Configure(machine.I2CConfig{SCL: machine.SCL_PIN, SDA: machine.SDA_PIN})

	// display := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	// display.Configure(st7735.Config{
	// 	Rotation: st7735.ROTATION_90,
	// })
	// display.FillScreen(color.RGBA{0, 0, 0, 255})

	// var data [64]int16
	// var value int16
	// for {
	// 	display.D

	// 	// get the values of the sensor in millicelsius
	// 	for j := int16(0); j < 8; j++ {
	// 		for i := int16(0); i < 8; i++ {
	// 			value = data[63-(i+j*8)]
	// 			// treat anything below 18°C as 18°C
	// 			if value < 18000 {
	// 				value = 0
	// 			} else {
	// 				value = (value - 18000) / 36
	// 				// our color array only have 433 values, avoid getting a value that doesn't exist
	// 				if value > 432 {
	// 					value = 432
	// 				}
	// 			}
	// 			// show the image on the PyBadge's display
	// 			display.FillRectangle(16+i*16, j*16, 16, 16, color.RGBA{240, 52, 52, 1})
	// 		}
	// 	}
	// }

}
