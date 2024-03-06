package server

import (
	"html/template"
	"log"
	"net/http"

	api "github.com/twelc/go-sheets/lib"
)

type Table struct {
	Data [][]string
}

var config = api.GetConfig("./config/credentials.json", "141maOrpeeFsydVAWP-kIaziMCHn_fI8nQv0mFB78TVk", "default")

func wrap(tpl *template.Template, err error) *template.Template {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return tpl
}

var tpl = wrap(template.ParseFiles("templates/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tab := api.GetAll(config, "A1:C500")
	tpl.Execute(w, Table{
		Data: tab,
	})
}
