package main

import (
	"flag"
	"fmt"
	"github.com/DrItanium/fakku"
	"log"
	"os"
	"text/tabwriter"
)

var category = flag.String("category", fakku.CategoryManga, "the type of the content")
var name = flag.String("name", "", "the name of the content itself, this is usually what you would find in the URL")
var downloadImages = flag.Bool("download", false, "Download the content you would read online to the current working directory")
var comments = flag.Bool("comments", false, "Show the content's comments")

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
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintf(tw, "Type:\t%s\n", *category)
	fmt.Fprintf(tw, "Name:\t%s\n", content.Name)
	fmt.Fprintf(tw, "Artists:\t%s\n", content.ArtistsString())
	fmt.Fprintf(tw, "Series:\t%s\n", content.SeriesString())
	fmt.Fprintf(tw, "Translators:\t%s\n", content.TranslatorsString())
	fmt.Fprintf(tw, "Tags:\t%s\n", content.TagsString())
	tw.Flush()

	if *comments {
		comments, errcm := content.Comments()
		if errcm != nil {
			log.Fatal(errcm)
		}
		fmt.Println("Comments")
		for _, comment := range comments.Comments {
			fmt.Printf("\t[%s] %s - %s\n", comment.Date(), comment.Poster, comment.Text)
		}
	}

	if *downloadImages {
		if derr0 := content.SaveCover("cover-thumbnail.jpg", 0644); derr0 != nil {
			log.Fatal(derr0)
		}
		// dump the read-online stuff
		pages, err1 := content.ReadOnline()
		if err1 != nil {
			log.Fatal(err1)
		}
		for ind, page := range pages {
			if dfErr := page.SaveImage(fmt.Sprintf("%03d.jpg", ind), 0644); dfErr != nil {
				log.Fatal(dfErr)
			}
		}
	}
}
