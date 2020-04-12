package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUser = map[string]user{}       //userid, user
var dbSessions = map[string]string{} //sessionid,user id

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {

	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{

			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c) //writing cookie into the session
	}

	// if user exists,get user

	var u user

	if un, ok := dbSessions[c.Value]; ok {
		u = dbUser[un]
	}

	// form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un //assigning user a uuid
		dbUser[un] = u
	}

	// if get method
	tpl.ExecuteTemplate(w, "index.html", u)

}

func bar(w http.ResponseWriter, r *http.Request) {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	// here !ok means there is no value stored
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	u := dbUser[un]
	tpl.ExecuteTemplate(w, "bar.html", u)

}
