package main

import (
	"fmt"
	"net/http"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Streaming content..."))
}

func main() {
	http.HandleFunc("/stream", streamHandler)
	fmt.Println("Streaming Service is running on port 8081")
	http.ListenAndServe(":8083", nil)
}
