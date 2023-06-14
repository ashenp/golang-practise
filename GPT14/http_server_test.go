package GPT14

import (
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	server := NewHttpServer(8080)
	server.RegisterFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	server.start()

}
