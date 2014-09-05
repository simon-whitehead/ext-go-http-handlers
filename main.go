package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", handler(myHandler))
	http.ListenAndServe(":8080", mux)
}

func myHandler(w http.ResponseWriter, r *http.Request, db *mgo.Database) {
	var users []user

	db.C("users").Find(nil).All(&users)

	for _, user := range users {
		fmt.Fprintf(w, "%s is %d years old", user.Name, user.Age)
	}
}
