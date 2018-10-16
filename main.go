package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/api/challenge", handleChallenge)
	appengine.Main()
}

func handleChallenge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var challenge challenge

	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(&challenge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "text/plain")
	w.Write([]byte(challenge.Challenge))
}

type challenge struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}
