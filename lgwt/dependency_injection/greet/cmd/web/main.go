package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zoumas/lab/lgwt/dependency_injection/greet"
)

func Greeter(name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		greet.Greet(w, name)
	})
}

func main() {
	const port = "8000"
	name := "World"

	fmt.Fprintln(os.Stderr, "Serving on port:", port)
	log.Fatalln(http.ListenAndServe(":"+port, Greeter(name)))
}
