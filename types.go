package main

// Note contains the metadata for a note
type Note struct {
	Title       string   `json:"title"`
	Attendees   []string `json:"attendees"`
	Created     string   `json:"created"`
	Path        string   `json:"path"`
	OutFile     string   `json:"out-file"`
	Template    string   `json:"template"`
	ActionItems []string `json:"action-items"`
}
