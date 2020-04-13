package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	// packages can be extended to different files in same folder
	// get cookie

	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{

			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	// if the user exists already, get user
	var u user
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUser[s.un]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUser[s.un]
	// refresh session
	c.MaxAge = sessionLength
	return ok

}

func cleanSessions() {
	fmt.Println("before clean") //demonstration use only
	// showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("After cleaned")
	// showSessions()
}

// for demonstation purposes
func showSessions() {
	fmt.Println("**********")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("**********")

}
