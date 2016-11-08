package main

import (
	"encoding/json"
	"github.com/tarm/serial"
	"log"
	"net/http"
	"strings"
)

var s *serial.Port

func main() {
	// create a config struct
	c := &serial.Config{
		Name: "/dev/tty.usbmodem1441",
		Baud: 9600,
	}

	// open the serial port using the config
	var err error
	s, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	// create a basic web server...
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	// we _should_ verify the request first -
	// https://docs.travis-ci.com/user/notifications/#Verifying-Webhook-requests

	// create data struct
	var d struct {
		Status  int    `json:"status"`
		Message string `json:"status_message"`
	}

	// decode payload
	err := json.NewDecoder(strings.NewReader(r.FormValue("payload"))).Decode(&d)
	if err != nil {
		log.Fatal(err)
	}

	// set color from status
	var c string
	switch d.Message {
	case "Pending":
		c = "255,255,0"
	case "Passed":
		c = "0,255,0"
	case "Failed":
		c = "255,0,0"
	}

	// write an r,g,b string to the serial port
	log.Print(c)

	w.Write([]byte("OK"))
}
