package main

import (
	"net/http"

	"github.com/mytkdals93/goweb/myFile"
)

func main() {
	http.ListenAndServe(":3000", myFile.NewHttpHanlder())
}
