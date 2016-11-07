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

	// write an r,g,b string to the serial pot
	n, err := s.Write([]byte("255,0,0\n"))
	if err != nil {
		log.Fatal(err)
	}

	// read any data that may come back from the serial port
	buf := make([]byte, 128)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	// print the data
	log.Printf("%q", buf[:n])

	// delay for a few seconds so we can see the color
	time.Sleep(time.Second * 2)

	// no need to close the serial, taum/serial does that for us
	// bounce 🏀
	return
}
