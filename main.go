package main 

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

type app struct {
	store map[string]string 
	mux *http.ServeMux
}

func newApp() *app {
	a := &app {
		store: make(map[string]string), 
		mux: http.NewServeMux(),
	}
	a.mux.HandleFunc("GET /shorten", a.handleShorten)
	a.mux.HandleFunc("GET /{code}", a.handleRedirect)
	return a 
}

// shorten function handler 
func (a *app) handleShorten(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "missing ?url=", http.StatusBadRequest)
		return 
	}
	
	// call our shorten function 
	code := shortCode()
	a.store[code] = url 
	fmt.Fprintf(w, "http://localhost:8080/%s\n", code) 


}


func (a *app) handleRedirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if url, ok := a.store[code]; ok {
		http.Redirect(w, r, url, http.StatusFound) 
	} else {
		http.NotFound(w, r)
	}
}



func shortCode() string {
	b := make([]byte, 4) 
	rand.Read(b) 
	return base64.URLEncoding.EncodeToString(b)[:6] 
}

func main() {
	a := newApp()
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", a.mux)
}

