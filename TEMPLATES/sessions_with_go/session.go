package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	// packages can be extended to different files in same folder

	// get cookie

	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c := &http.Cookie{

			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, c)
	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUser[un]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUser[un]
	return ok

}
