package main

import (
	"net/http"

	"web5/app"
)

func main() {
	http.ListenAndServe(":3000", app.NewHanlder())
}
