package main

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var err error

type handler func(w http.ResponseWriter, r *http.Request, db *mgo.Database)

func init() {
	session, err = mgo.Dial("localhost")

	if err != nil {
		log.Println(err)
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := session.Clone()
	defer s.Close()

	h(w, r, s.DB("example"))
}
