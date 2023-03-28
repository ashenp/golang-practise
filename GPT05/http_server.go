package GPT05

import (
	"fmt"
	"net/http"
)

func HttpServer() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		query := r.URL.Query()
		name := query.Get("name")
		fmt.Fprintf(w, "Hello, %s!", name)
	}

	if r.Method == http.MethodPost {
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		fmt.Fprintf(w, "Got body: %q", body)
	}
}
