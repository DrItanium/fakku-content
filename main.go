package main

import (
	"flag"
	"fmt"
	"github.com/DrItanium/fakku"
	"log"
)

var category = flag.String("category", fakku.Manga, "the type of the content")
var name = flag.String("name", "", "the name of the content itself, this is usually what you would find in the URL")

func main() {
	flag.Parse()
	if *name == "" {
		log.Fatal("Did not provide a name")
	}
	if !fakku.LegalCategory(*category) {
		log.Fatalf("Illegal category %s provided", *category)
	}
	content, err := fakku.GetContent(*category, *name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name:", content.Name)
}
