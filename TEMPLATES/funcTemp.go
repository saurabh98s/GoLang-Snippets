package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var (
	tpl *template.Template
)
var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {

	return t.Format("01-02-2006")
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
