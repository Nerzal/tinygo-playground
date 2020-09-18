package main

import (
	"device/avr"
	"machine"
	"time"
)

const (
	Mode0 uint8 = iota
	Mode1
	Mode2
	Mode3
)

type SPIClock uint8

const (
	SPI_CLOCK_FCK2   SPIClock = 0
	SPI_CLOCK_FCK4            = 1
	SPI_CLOCK_FCK8            = 2
	SPI_CLOCK_FCK16           = 3
	SPI_CLOCK_FCK32           = 4
	SPI_CLOCK_FCK64           = 5
	SPI_CLOCK_FCK128          = 7

	SPI_CLOCK_MASK   = 0x03
	SPI_2XCLOCK_MASK = 0x01
)

type SPIConfig struct {
	IsSlave  bool
	LSB      bool
	MaxSpeed SPIClock
	Mode     uint8
	SDI      machine.Pin
	SDO      machine.Pin
	SCK      machine.Pin
}

type SPI struct {
}

var SPI0 = SPI{}

func (spi SPI) Configure(config SPIConfig) {
	setMode(config.Mode)

	if config.LSB {
		spi.LSB()
	} else {
		spi.MSB()
	}

	// Invert the SPI2X bit
	config.MaxSpeed ^= 0x1
	avr.SPSR.SetBits(uint8(config.MaxSpeed) & SPI_2XCLOCK_MASK)

	if config.IsSlave {
		spi.Slave(config)
	} else {
		spi.Master(config)
	}
}

func (s SPI) LSB() {
	avr.SPCR.SetBits(avr.SPCR_DORD)
}

func (s SPI) MSB() {
	avr.SPCR.ClearBits(avr.SPCR_DORD)
}

func (spi SPI) Master(config SPIConfig) {
	avr.DDRB.SetBits(uint8(config.SDO) | uint8(config.SCK))        // set sdo, sck as output, all other input
	avr.DDRB.ClearBits(1 << 4)                                     // sck is high when idle
	avr.SPCR.SetBits(avr.SPCR_MSTR | avr.SPCR_SPR0 | avr.SPCR_SPE) // set master, set clock rate fck/16, enable spi
}

func (spi SPI) Slave(s SPIConfig) {
	avr.DDRB.SetBits(1 << uint8(s.SDI))            // set sdi output, all other input
	avr.SPCR.ClearBits(avr.SPCR_MSTR)              // set slave
	avr.SPCR.SetBits(avr.SPCR_SPR0 | avr.SPCR_SPE) // set clock rate fck/16, enable spi
}

func (SPI) Transfer(b byte) byte {
	avr.SPDR.Set(uint8(b))

	waitForRegisterShift()

	return byte(avr.SPDR.Reg)
}

func (s SPI) Receive() byte {
	waitForRegisterShift()

	return byte(avr.SPDR.Reg)
}

func (s SPI) Send(b byte) {
	avr.SPDR.Set(uint8(b))

	waitForRegisterShift()
}

func waitForRegisterShift() {
	for !avr.SPSR.HasBits(avr.SPSR_SPIF) {
	}
}

func setMode(mode uint8) {
	switch mode {
	case 0:
		avr.SPCR.ClearBits(avr.SPCR_CPOL)
		avr.SPCR.ClearBits(avr.SPCR_CPHA)
	case 1:
		avr.SPCR.ClearBits(avr.SPCR_CPOL)
		avr.SPCR.SetBits(avr.SPCR_CPHA)
	case 2:
		avr.SPCR.SetBits(avr.SPCR_CPOL)
		avr.SPCR.ClearBits(avr.SPCR_CPHA)
	case 3:
		avr.SPCR.SetBits(avr.SPCR_CPOL)
		avr.SPCR.SetBits(avr.SPCR_CPHA)
	default:
		avr.SPCR.ClearBits(avr.SPCR_CPOL)
		avr.SPCR.ClearBits(avr.SPCR_CPHA)
	}
}

func init() {
	println("initializing")
	SPI0.Configure(SPIConfig{
		Mode:     0,
		SDI:      machine.PB4,
		SDO:      machine.PB3,
		SCK:      machine.PB5,
		LSB:      true,
		MaxSpeed: SPI_CLOCK_FCK4,
	})
}

func main() {
	for {
		println("loopedi loop loop")
		for a := byte('!'); a < 127; a++ {
			print("input: ", a, " ")
			b := SPI0.Transfer(a)
			println("data read: ", b)
			time.Sleep(1 * time.Second)
		}
	}
}
