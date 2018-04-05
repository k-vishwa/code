package main

import (
	"fmt"
	"net/http"
)

var HelloMessage string

func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	out, _ := exec.Command("bash", "-c", "hostname").Output()
// 	HelloMessage = "#############" + string(out)
	HelloMessage = "Hello World-canary"
  	fmt.Fprintf(w, HelloMessage)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8081", nil)
}
