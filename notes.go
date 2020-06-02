package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

// NewNote will return a note metadata object
func NewNote(title *string, attendees *string, path *string, template *string) (*Note, error) {
	note := Note{
		Title:     *title,
		Attendees: strings.Split(*attendees, " "),
		Path:      *path,
		Template:  *template,
	}
	now := time.Now()
	note.Created = now.Format("2006-01-02")

	if note.Path == "" {
		note.Path = fmt.Sprintf("%s-%s.md", note.Created, strings.Replace(note.Title, " ", "-", -1))
	}

	return &note, nil
}

// CreateNote will write the note template to disk
func (n *Note) CreateNote() error {
	tmpl, err := template.ParseFiles("templates/one-on-one.md.j2")
	if err != nil {
		return err
	}

	f, err := os.Create(n.Path)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(f, n)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}
