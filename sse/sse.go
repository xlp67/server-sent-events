package sse

import (
	"fmt"
	"net/http"
	"sse/file"
)

func ResponseSSE(filePath string, w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") != "text/event-stream" {
		http.Error(w, "O cliente não suporta SSE", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	bytesData := file.ReadFile(filePath)

	for i := 0; i < len(bytesData); i ++ {
		fmt.Fprintf(w, "data: %x\r\n\r\n", bytesData[i])
		w.(http.Flusher).Flush()
	}
}
