package myFile

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err2 := os.Create(filepath)
	defer file.Close()
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err2)
		return
	}
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}
func NewHttpHanlder() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(`./myFile/public`)))
	mux.HandleFunc("/uploads", uploadsHandler)
	return mux
}
