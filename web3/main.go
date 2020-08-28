package main

import (
	"net/http"

	"github.com/mytkdals93/goweb/web3/app"
)

func main() {
	http.ListenAndServe(":3000", app.NewHTTPHanlder())
}
