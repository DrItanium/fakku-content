package main

import (
	"flag"
	"github.com/DrItanium/fakku"
	"log"
)

var group = flag.String("group", fakku.Manga, "the type of the content")
var name = flag.String("name", "", "the name of the content itself, this is usually what you would find in the URL")

func main() {
	flag.Parse()
	if *name == "" {
		log.Fatal("Did not provide a name")
	}
	if !fakku.LegalGroup(*group) {
		log.Fatalf("Illegal group %s provided", *group)
	}

}
