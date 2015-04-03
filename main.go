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
	covUrl, errcu := content.CoverUrl()
	if errcu != nil {
		log.Fatal(errcu)
	}
	fmt.Println("Cover URL:", covUrl.String())
	comments, errcm := content.Comments()
	if errcm != nil {
		log.Fatal(errcm)
	}
	fmt.Println("Comments")
	for _, comment := range comments.Comments {
		fmt.Printf("[%s] %s - %s\n", comment.Date(), comment.Poster, comment.Text)
	}
}
