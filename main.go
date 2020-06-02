package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// GitCommit provided at build time
var GitCommit string

// GitState provided at build time
var GitState string

// Version provided at build time
var Version string

func main() {

	required := []string{"title"}
	seen := make(map[string]bool)
	titlePtr := flag.String("title", "", "(requried) Title of note")
	attendeesPtr := flag.String("attendees", "", "(optional) List of Attendees")
	outPtr := flag.String("out", "", "(optional) Output file name")
	pathPtr := flag.String("path", "./", "(optional) Path to new note. Defaults to current directory")

	flag.Parse()

	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			fmt.Fprintf(os.Stderr, "missing required -%s argument/flag\n\nUsage:\n", req)

			flag.PrintDefaults()
			os.Exit(2)
		}
	}
	templ := "one-on-one.md.j2"
	note, err := NewNote(titlePtr, attendeesPtr, outPtr, pathPtr, &templ)
	if err != nil {
		log.Fatal(err)
	}

	err = note.CreateNote()
	if err != nil {
		log.Fatal(err)
	}

}
