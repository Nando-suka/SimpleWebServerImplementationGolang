package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Cleaner, conventional naming
type ClubInfo struct {
	ClubName string
	City     string
}

type Personnn struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    ClubInfo
}

type AppState struct {
	CurrentStatus string
	StateMessage  string
}

func main() {
	appState := AppState{
		CurrentStatus: "Server is running",
		StateMessage:  "All systems nominal",
	}

	// Parse templates once at startup, not on every request
	tmpl, err := template.ParseFiles("viewW.html")
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		people := []Personnn{
			{
				Name:    "Erling Braut Haaland",
				Gender:  "Male",
				Hobbies: []string{"Playing Football", "Social Media", "Being funny for teammates"},
				Info:    ClubInfo{"Manchester City FC", "Manchester, England"},
			},
			{
				Name:    "Pedro Neto",
				Gender:  "Male",
				Hobbies: []string{"Playing Football", "Hanging with friends", "Teasing Cole Palmer"},
				Info:    ClubInfo{"Chelsea FC", "London, England"},
			},
		}

		// Set content type explicitly
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err := tmpl.Execute(w, people); err != nil {
			// Log the real error server-side, send generic message to client
			log.Printf("Template execution error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Printf("App state: %+v\n", appState)
	})

	http.HandleFunc("/aboutMe", func(w http.ResponseWriter, r *http.Request) {
		// This was empty — add a basic response so it doesn't silently do nothing
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "About Me page — coming soon!")
	})

	addr := ":9070"
	fmt.Printf("Server started at http://localhost%s\n", addr)

	// ListenAndServe always returns an error — you should handle it
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
