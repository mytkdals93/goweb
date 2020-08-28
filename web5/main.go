package main

import (
	"net/http"

	"github.com/mytkdals93/goweb/web5/app"
)

func main() {
	http.ListenAndServe(":3000", app.NewHanlder())
}
