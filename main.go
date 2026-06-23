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
	
	// 1. API Routes
	a.mux.HandleFunc("GET /shorten", a.handleShorten)
	a.mux.HandleFunc("GET /{code}", a.handleRedirect)
	
	// 2. Frontend Static Files
	// We serve them under "/frontend/" to avoid conflicting with "/{code}"
	fileServer := http.FileServer(http.Dir("./frontend"))
	a.mux.Handle("GET /frontend/", http.StripPrefix("/frontend", fileServer))
	
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
	
	// Return the short URL (removed the \n so it's cleaner for the frontend)
	fmt.Fprintf(w, "http://localhost:8080/%s", code) 
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
	fmt.Println("Frontend available at: http://localhost:8080/frontend/")
	http.ListenAndServe(":8080", a.mux)
}