package main

import (
	"fmt"
	"net/http"
)

var HelloMessage string

func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	out, _ := exec.Command("bash", "-c", "hostname").Output()
// 	HelloMessage = "#############" + string(out)
	HelloMessage = "Hello World-bg1"
	fmt.Println(HelloMessage)
  	fmt.Fprintf(w, HelloMessage)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":4444", nil)
}
