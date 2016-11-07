package main

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

func main() {
	// create a config struct
	c := &serial.Config{
		Name: "/dev/tty.usbmodem1441",
		Baud: 9600,
	}

	// open the serial port using the config
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	// write an r,g,b string to the serial port
	_, err = s.Write([]byte("255,0,0\n"))
	if err != nil {
		log.Fatal(err)
	}

	// delay for a few seconds so we can see the color
	time.Sleep(time.Second * 1)

	// write another r,g,b string to the serial port
	_, err = s.Write([]byte("0,255,0\n"))
	if err != nil {
		log.Fatal(err)
	}

	// delay for a few seconds so we can see the color
	time.Sleep(time.Second * 1)

	// write another r,g,b string to the serial port
	_, err = s.Write([]byte("0,0,255\n"))
	if err != nil {
		log.Fatal(err)
	}

	// delay for a few seconds so we can see the color
	time.Sleep(time.Second * 1)
}
