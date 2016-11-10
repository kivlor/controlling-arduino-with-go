package main

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

func main() {
	// create a config struct
	c := &serial.Config{
		Name: "/dev/tty.usbmodem1234", // Change this to your device
		Baud: 9600,                    // Change this to your devices data rate for serial
	}

	// open the serial port using the config
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	// infinate loop!
	for {
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
}
