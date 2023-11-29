package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func ResponseSSE(channel chan int, w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") != "text/event-stream" {
		http.Error(w, "O cliente não suporta SSE", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for i := range channel {
		message := fmt.Sprintf("%d", i)
		fmt.Fprintf(w, "data: %s\n\n", message)
		w.(http.Flusher).Flush()
	}
}

func ProcessData(channel chan int, data int) {
	for i := 0; i < data; i++ {
		channel <- i
	}
	close(channel)
}

func Template(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {log.Fatal("erro no parse no template")}
	tpl.Execute(w, nil)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/partial-response", func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan int)
		go ResponseSSE(ch, w, r)
		ProcessData(ch, 10)
	})
	mux.HandleFunc("/", Template)
	http.ListenAndServe(":8080", mux)
}