package main

import (
	"net/http"
)

func IndexHandler (w http.ResponseWriter, r *http.Request) {
	vars := PageVars{
		Title: "Golang",
	}
	render(w, "index.html", vars);
}