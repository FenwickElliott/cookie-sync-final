package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var port = flag.String("port", "4000", "port to serve on")

func main() {
	http.HandleFunc("/", root)
	fmt.Println("Serving on port:", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("glamID")
	if cookie == nil {
		cookie = setCookie(&w, r)
	} else {
		check(err)
	}
}

// Utility functions:
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func setCookie(w *http.ResponseWriter, r *http.Request) *http.Cookie {
	h := sha1.New()
	h.Write([]byte(time.Now().String() + r.RemoteAddr))
	cookie := http.Cookie{Name: "glamID", Value: hex.EncodeToString(h.Sum(nil)), Expires: time.Now().Add(365 * 24 * time.Hour)}
	http.SetCookie(*w, &cookie)
	return &cookie
}
