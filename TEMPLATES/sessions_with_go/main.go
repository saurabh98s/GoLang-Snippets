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
	// to implement the logic logic we need to have a pre-existing user
	// since there is no DB connection yet,we will create one and store it in the map
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUser["test@test.com"] = user{"test@test.com", bs, "James", "Bond"}
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.html", u)

}

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

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		u, ok := dbUser[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusSeeOther)
			return
		}

		//create session

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", nil)

}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	// delete the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
