package main

import (
	"encoding/json"
	"fmt"
	"github.com/tarm/serial"
	"html/template"
	"log"
	"net/http"
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

	// crate a basic web server...
	http.HandleFunc("/", root)
	http.HandleFunc("/update", update)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("home").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>color picker</title>
    <style>
      * { box-sizing: border-box; }
      div#colors { margin: 20px auto; width: 600px; overflow: hidden; }
      div#colors div { display: inline; float: left; width: 200px; height: 200px; border: 10px solid #fff; cursor: pointer; }
    </style>
  </head>
  <body>
    <div id="colors">
      <div style="background: rgb(255, 0, 0);" onclick="sendColor(255, 0, 0)"></div>
      <div style="background: rgb(255, 255, 0);" onclick="sendColor(255, 255, 0)"></div>
      <div style="background: rgb(0, 255, 0);" onclick="sendColor(0, 255, 0)"></div>
      <div style="background: rgb(0, 255, 255);" onclick="sendColor(0, 255, 255)"></div>
      <div style="background: rgb(0, 0, 255);" onclick="sendColor(0, 0, 255)"></div>
      <div style="background: rgb(255, 255, 255);" onclick="sendColor(255, 255, 255)"></div>
    </div>
    <script type="text/javascript">
      function sendColor(red, green, blue) {
        var request = new XMLHttpRequest();
        request.open('POST', '/update', true);
        request.setRequestHeader('Content-Type', 'application/json');
        request.send(JSON.stringify({ r: +red, g: +green, b: +blue }));
      }
    </script>
  </body>
</html>
`)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func update(w http.ResponseWriter, r *http.Request) {
	// create color struct
	var c struct {
		Red   int `json:"r"`
		Green int `json:"g"`
		Blue  int `json:"b"`
	}

	// decode json data
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	// write an r,g,b string to the serial port
	_, err = s.Write([]byte(fmt.Sprintf("%v,%v,%v\n", c.Red, c.Green, c.Blue)))
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte("OK"))
}
