package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func SSE(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") != "text/event-stream" {
		http.Error(w, "O cliente não suporta SSE", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	channel := make(chan string)
	defer close(channel)

	for i := 0; i < 100; i++ {
		message := fmt.Sprintf("%d", i)
		fmt.Fprintf(w, "data: %s\n\n", message)
		w.(http.Flusher).Flush()
		time.Sleep(time.Second) 
	}
}
func Template(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {log.Fatal("erro no parse no template")}
	tpl.Execute(w, nil)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/partial-response", SSE)
	mux.HandleFunc("/", Template)
	http.ListenAndServe(":8080", mux)
}