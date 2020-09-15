module github.com/Nerzal/tinygo-playground

go 1.14

replace machine => /home/tobias/go/src/github.com/tinygo-org/tinygo/src/machine

replace runtime/volatile => /home/tobias/go/src/github.com/tinygo-org/tinygo/src/runtime/volatile

require (
	machine v0.0.0-00010101000000-000000000000
	runtime/volatile v0.0.0-00010101000000-000000000000
	tinygo.org/x/drivers v0.13.0
)
