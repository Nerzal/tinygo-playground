# tinygo-playground

Examples do work with Arduino Uno

## Flashing blink-sos

>  tinygo flash --target=arduino cmd/blink-sos/sos.go

## Flashing traffic-light

>  tinygo flash --target=arduino cmd/traffic-light/light.go

## Common Problems

> avrdude: ser_open(): can't open device "/dev/ttyACM0": Permission denied

Change persmissions of the serial port by using the following command
> sudo chmod a+rw /dev/ttyACM0