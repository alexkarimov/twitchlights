package main

//This is merely a template of what I use
//If you wish to use this for yourself you'll need to fill in the "hueURL" and "readcolour" with your own values.
//Original version made by alexkarimov on github.com
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var currentColor string

//Fill in your local connection to your bridge here
const hueURL = "http://192.168.0.22/api/---FILL-IN-BRIDGE-LOGIN-INFO-HERE---/groups/0/action"

// readColor - Reads colour from the server.
func readColor() string {
	resp, err := http.Get("http://---API-LINK-HERE---.appspot.com/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

//These are colours I created on my phone that I pulled from my computer and pasted here
var hueColors = map[string]string{
	"red":    "f30d50551-on-0",
	"blue":   "af32ef627-on-0",
	"purple": "759769410-on-0",
	"pink":   "3c658c721-on-0",
	"yellow": "7f3f1cb41-on-0",
	"green":  "13f0a3309-on-0",
	"orange": "99e79a27c-on-0",
	"white":  "80cb19250-on-0",
	"black":  "10f591d1d-on-0",
	"brown":  "de36c6900-on-0",
}

// changeColor sends info the hue station on Wifi
func changeColor(color string) {
	fmt.Println("changing colour to ", color)

	value := hueColors[color]
	if value == "" {
		fmt.Println("Unknown color:", color)
		return
	}

	message := fmt.Sprintf(`{"scene": "%s"}`, value)

	req, err := http.NewRequest("PUT", hueURL, bytes.NewBufferString(message))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}

//If colour is the same (On API and local), don't do anything
//If colour value changes on API, change local
func main() {
	for {
		color := readColor()

		if currentColor != color {
			changeColor(color)
			currentColor = color
		}
		//You can change this number to whatever you like, I personally have it default set to this.
		time.Sleep(500 * time.Millisecond)
	}
}
