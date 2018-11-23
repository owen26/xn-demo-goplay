package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dchest/uniuri"
	"github.com/satori/go.uuid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	u1 := uuid.NewV4()
	u2 := uniuri.New()
	fmt.Fprintf(w, "Hi there, I love %s! %s %s", r.URL.Path[1:], u1, u2)
}

func tomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Tom, I love %s!", r.URL.Path[1:])
}

func jerryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Jerry, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/tom", tomHandler)
	http.HandleFunc("/jerry", jerryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
