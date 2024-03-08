package server

import (
	"encoding/json"
	"html/template"
	"net/http"

	api "github.com/twelc/go-sheets/lib"
)

var graphTpl = wrap(template.ParseFiles("templates/index.html"))

type querry struct {
	Querry string
}

func graphHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		q := r.FormValue("querry")
		graphTpl.Execute(w, querry{
			Querry: q,
		})
	}
}

func getGraphHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		querry := r.FormValue("querry")
		coupert := r.FormValue("coupert")
		conf := api.GetConfig(credentials, sheetid, coupert)
		api.NewSheet(conf)
		res, datas := api.GetCalculatedGraphData(querry, "history", "A:B", 3500, conf)
		defer api.DeleteSheet(conf)
		w.Header().Set("Content-Type", "application/json")
		resp := Response{
			Data: res,
			Time: datas,
		}

		json.NewEncoder(w).Encode(resp)
	}
}
