package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.String("port", "3000", "port to serve on")

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	fmt.Println("Serving on port:", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
