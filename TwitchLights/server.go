package hello

import {
"fmt"
"net/http"
}

//Handlers get added here
//Make sure one exists for every colour that you plan on using
func init() {
http.HandleFunc("/", mainhandler)
http.HandleFunc("/setblue", setblue)
http.HandleFunc("/setred", setred)
http.HandleFunc("/setgreen", setgreen)
}

//Colors get added here
var color string
func setblue(w http.ResponseWriter, r *http.Request) {
color = "blue"
//The "Fprint()" will be what Nightbot (or any other chatbot going to the API) will send back into twitchchat.
//You can set whatever you want the text to be in the quotes.
fmt.Fprint(w, "Setting to blue!")
}

//Below are two more examples just to show how you can add more colours
func setred(w http.ResponseWriter, r *http.Request) {
color = "red"
fmt.Fprint(w, "RED! RED! RED! RED")
}
func setgreen(w http.ResponseWriter, r *http.Request) {
color = "green"
fmt.Fprint(w, "Wowee! You've set it to green!")
}

//Main Page
func mainhandler(w http.ResponseWriter, r *http.Request) {
fmt.Fprint(w, color)
}
