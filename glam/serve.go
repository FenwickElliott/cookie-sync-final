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
	http.HandleFunc("/in", in)
	fmt.Println("Serving on port:", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func in(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	partnerID := r.FormValue("partner")
	partnerCookie := r.FormValue("cookieID")

	glamCookie, err := r.Cookie("glamID")
	if glamCookie == nil {
		glamCookie = setCookie(&w, r)
	} else {
		check(err)
	}

	fmt.Println("glamID:", glamCookie.Value)
	fmt.Println("partnerID:", partnerID)
	fmt.Println("partnerCookie:", partnerCookie)
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
