package server

import (
	"net/http"
	"os"
)

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/get-table-data", getTableHandler)
	mux.HandleFunc("/get-table-filter", filteredHandler)
	mux.HandleFunc("/get-graph-data", getGraphHandler)
	mux.HandleFunc("/graph", graphHandler)
	http.ListenAndServe(":"+port, mux)
}
