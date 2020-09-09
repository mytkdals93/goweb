package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/antage/eventsource"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	msg := new(Message)
	err := json.NewDecoder(r.Body).Decode(msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	log.Println("postMessageHandler", msg.Msg)
	sendMessage(msg)
	data, _ := json.Marshal(msg)
	fmt.Fprint(w, string(data))
}
func addUserHandler(w http.ResponseWriter, r *http.Request) {
	msg := new(Message)
	err := json.NewDecoder(r.Body).Decode(msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	sendMessage(&Message{"", fmt.Sprintf("add user: %s", msg.Name)})
}
func leftUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	sendMessage(&Message{"", fmt.Sprintf("Left user: %s", username)})
}

type Message struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

var msgCh chan Message

func sendMessage(msg *Message) {
	//send message to every client
	msgCh <- *msg
}
func processMsgCh(es eventsource.EventSource) {
	for msg := range msgCh {
		data, _ := json.Marshal(msg)
		es.SendEventMessage(string(data), "", strconv.Itoa(time.Now().Nanosecond()))
	}
}
func main() {
	msgCh = make(chan Message)
	es := eventsource.New(nil, nil)
	defer es.Close()

	go processMsgCh(es)

	mux := pat.New()
	mux.Post("/messages", postMessageHandler)
	mux.Handle("/stream", es)
	mux.Delete("/users", leftUserHandler)
	mux.Post("/users", addUserHandler)
	n := negroni.Classic()
	n.UseHandler(mux)
	log.Fatal(http.ListenAndServe(":3000", n))
}
