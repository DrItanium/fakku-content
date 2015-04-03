package main

import (
	"flag"
	"fmt"
	"github.com/DrItanium/fakku"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
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
		fmt.Printf("\t[%s] %s - %s\n", comment.Date(), comment.Poster, comment.Text)
	}

	// try getting the covUrl's contents
	derr0 := DownloadFile(covUrl, "cover-thumbnail.jpg", 0644)
	if derr0 != nil {
		log.Fatal(derr0)
	}
	// dump the read-online stuff
	pages, err1 := content.ReadOnline()
	if err1 != nil {
		log.Fatal(err1)
	}
	for ind, page := range pages {
		purl, perr := page.ImageUrl()
		if perr != nil {
			log.Fatal(perr)
		}
		dfErr := DownloadFile(purl, fmt.Sprintf("%03d.jpg", ind), 0644)
		if dfErr != nil {
			log.Print(dfErr)
		}
	}
}

func DownloadFile(url *url.URL, outputDir string, perms os.FileMode) error {
	resp, rerr := http.Get(url.String())
	if rerr != nil {
		return rerr
	}
	defer resp.Body.Close()
	img, ierr := ioutil.ReadAll(resp.Body)

	if ierr != nil {
		return ierr
	}
	werr := ioutil.WriteFile(outputDir, img, perms)
	if werr != nil {
		return werr
	}
	return nil
}
