package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func ResponseSSE(channel chan rune, w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") != "text/event-stream" {
		http.Error(w, "O cliente não suporta SSE", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for char := range channel {
		message := fmt.Sprintf("%c", char)
		fmt.Fprintf(w, "data: %s\n\n", message)
		w.(http.Flusher).Flush()
	}
}

func ReadFile(channel chan rune, path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, char := range string(content) {
		channel <- char
	}
	close(channel)
}

func Template(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal("erro no parse no template")
	}
	tpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/partial-response", func(w http.ResponseWriter, r *http.Request) {
		ch := make(chan rune)
		go ResponseSSE(ch, w, r)
		ReadFile(ch, "../resolve-ai/main.go")
	})
	mux.HandleFunc("/", Template)
	http.ListenAndServe(":8080", mux)
}
