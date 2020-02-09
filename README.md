# tinygo-playground

## Common Problems

> avrdude: ser_open(): can't open device "/dev/ttyACM0": Permission denied

Change persmissions of the serial port by using the following command
> sudo chmod a+rw /dev/ttyACM0