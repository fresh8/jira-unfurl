package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/api/events", handleEvents)
	appengine.Main()
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var event Event

	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)

	err := dec.Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch event.Type {
	case "url_verification":
		w.Header().Set("Content-type", "text/plain")
		w.Write([]byte(event.Challenge))
		return

	}
}

type Event struct {
	Token       string   `json:"token"`
	Challenge   string   `json:"challenge,omitempty"`
	Type        string   `json:"type"`
	TeamID      string   `json:"team_id"`
	ApiAppID    string   `json:"api_app_id"`
	AuthedUsers []string `json:"authed_users"`
	EventID     string   `json:"event_id"`
}
