package server

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	api "github.com/twelc/go-sheets/lib"
)

var graphTpl = wrap(template.ParseFiles("templates/graph.html"))

type querry struct {
	Querry string
}

func graphHandler(w http.ResponseWriter, r *http.Request) {
	q := strings.Replace(r.URL.Path, "/graph/", "", -1)
	graphTpl.Execute(w, querry{
		Querry: q,
	})
}

func getRandHash() string {
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

func getGraphHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("responcing")
		r.ParseForm()
		querry := r.FormValue("querry")
		hash := getRandHash()
		conf := api.GetConfig(credentials, sheetid, hash)
		api.NewSheet(conf)
		res, datas := api.GetCalculatedGraphData(querry, "history", "A:B", 3500, conf)
		//defer api.DeleteSheet(conf)
		w.Header().Set("Content-Type", "application/json")
		resp := Response{
			Data: res,
			Time: datas,
		}

		json.NewEncoder(w).Encode(resp)
	}
}
