package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

var kvalidPath = regexp.MustCompile("^/(demoEndPoint)/$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		m := kvalidPath.FindStringSubmatch(request.URL.Path)
		if m == nil {
			http.NotFound(writer, request)
			return
		}
		fn(writer, request, m[1])
	}
}

func genericHandler(writer http.ResponseWriter, request *http.Request, title string) {
	// fun fun fun fun
}

func demoEndPointHandler(writer http.ResponseWriter, request *http.Request, title string) {
	fmt.Fprintf(writer, "<html><h1>Brady is the best</h1><div><img src=%s></div>", "https://i.ytimg.com/vi/3EIbWjkimAs/maxresdefault.jpg")
}

func main() {
	http.HandleFunc("/", makeHandler(genericHandler))
	http.HandleFunc("/demoEndPoint/", makeHandler(demoEndPointHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Successfully started server on address: %s, port #: %s\n", "localhost", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}