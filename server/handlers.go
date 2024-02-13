package server

import (
	"html/template"
	"log"
	"net/http"

	api "github.com/twelc/go-sheets/api"
)

type Table struct {
	Data [][]string
}

func wrap(tpl *template.Template, err error) *template.Template {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return tpl
}

var tpl = wrap(template.ParseFiles("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tab := api.GetAll()
	tpl.Execute(w, Table{
		Data: tab,
	})
}
