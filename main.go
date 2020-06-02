package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

type note struct {
	Title     string `json:"title"`
	Attendees string `json:"attendees"`
}

// GitCommit provided at build time
var GitCommit string

// GitState provided at build time
var GitState string

// Version provided at build time
var Version string

func main() {

	titlePtr := flag.String("title", "", "Title")
	attendees := flag.String("attendees", "", "Attendees")

	if *titlePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	data := note{
		Title:     *titlePtr,
		Attendees: *attendees,
	}

	tmpl, err := template.ParseFiles("templates/one-on-one.md.j2")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("out.md")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}
