package main

import (
	"net/http"

	"github.com/mytkdals93/goweb/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHanlder())
}
