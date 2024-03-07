package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

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
var table = wrap(template.ParseFiles("templates/table.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tab := api.GetAll(config, "A1:C500")
	tpl.Execute(w, Table{
		Data: tab,
	})
}

func getTableHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		range_ := r.FormValue("range")
		ind, _ := strconv.Atoi(range_)
		fmt.Println(ind)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		table_range := fmt.Sprintf("A%v:C%v", ind*500+1, (ind+1)*500)
		tab := api.GetAll(config, table_range)

		if len(tab) == 0 {
			return
		} else {
			table.Execute(w, Table{
				Data: tab,
			})
		}
	}
}

func filteredHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		querry := r.FormValue("querry")
		min := r.FormValue("min")
		max := r.FormValue("max")
		range_ := r.FormValue("range")
		ind, _ := strconv.Atoi(range_)
		tab := api.GetFiltered(querry, min, max, "A1:C5000", config)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if len(tab) == 0 {
			return
		} else if len(tab) > 600 {
			tab = tab[ind*500 : (ind+1)*500]
			table.Execute(w, Table{
				Data: tab,
			})
		} else {
			table.Execute(w, Table{
				Data: tab,
			})
		}
	}
}
