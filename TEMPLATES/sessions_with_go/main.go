package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	Password []byte
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
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.html", u)

}

// func foo(w http.ResponseWriter, r *http.Request) {

// 	// get cookie
// 	c, err := r.Cookie("session")
// 	if err != nil {
// 		sID, _ := uuid.NewV4()
// 		c = &http.Cookie{

// 			Name:  "session",
// 			Value: sID.String(),
// 		}
// 		http.SetCookie(w, c) //writing cookie into the session
// 	}

// 	// if user exists,get user

// 	var u user

// 	if un, ok := dbSessions[c.Value]; ok {
// 		u = dbUser[un]
// 	}

// 	// form submission
// 	if r.Method == http.MethodPost {
// 		un := r.FormValue("username")
// 		p := r.FormValue("password")
// 		f := r.FormValue("firstname")
// 		l := r.FormValue("lastname")
// 		u = user{un, p, f, l}
// 		dbSessions[c.Value] = un //assigning user a uuid
// 		dbUser[un] = u
// 	}

// 	// if get method
// 	tpl.ExecuteTemplate(w, "index.html", u)

// }

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.html", u)

}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// process form submission

	if r.Method == http.MethodPost {
		// get form values
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")

		// username taken?
		if _, ok := dbUser[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return

		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		u := user{un, bs, f, l}
		dbUser[un] = u

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)

}
