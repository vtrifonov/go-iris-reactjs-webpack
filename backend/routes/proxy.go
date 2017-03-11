package routes

import (
	"io"
	"log"
	"net/http"
)

func ReloadProxy(w http.ResponseWriter, r *http.Request) {
	log.Println("Debug, Hot reload", r.Host)
	resp, err := http.Get("http://localhost:3000" + r.RequestURI)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
