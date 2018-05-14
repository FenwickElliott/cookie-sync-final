package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	port        = flag.String("port", "5000", "port to serve on")
	mongoServer = flag.String("mongoServer", "127.0.0.1", "mongo server address")

	db  *mgo.Database
	err error
)

type association struct {
	RockID        string
	PartnerCookie string
}

func main() {
	session, err := mgo.Dial(*mongoServer)
	check(err)
	defer session.Close()
	db = session.DB("rock")

	http.HandleFunc("/in", in)
	fmt.Println("Serving on port:", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func in(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	partnerID := r.FormValue("partner")
	partnerCookie := r.FormValue("cookieID")

	rockCookie, err := r.Cookie("rockID")
	if rockCookie == nil {
		rockCookie = setCookie(&w, r)
	} else {
		check(err)
	}

	res := association{}
	c := db.C(partnerID)
	err = c.Find(bson.M{"rockid": rockCookie.Value}).One(&res)
	if err != nil {
		c.Insert(association{rockCookie.Value, partnerCookie})
		err = c.Find(bson.M{"rockid": rockCookie.Value}).One(&res)
	}
	check(err)
	if res.PartnerCookie != partnerCookie {
		panic("partnerCookie doesn't match")
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
	cookie := http.Cookie{Name: "rockID", Value: hex.EncodeToString(h.Sum(nil)), Expires: time.Now().Add(365 * 24 * time.Hour)}
	http.SetCookie(*w, &cookie)
	return &cookie
}
