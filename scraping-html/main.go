package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Quote struct {
	Text   string
	Author string
	Tags   []string
}

func main() {
	res, err := http.Get("https://quotes.toscrape.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	rows := make([]Quote, 0)

	doc.Find(".col-md-8").ChildrenFiltered(".quote").Each(func(i int, sel *goquery.Selection) {
		row := new(Quote)
		row.Text = sel.Find(".text").Text()
		row.Author = sel.Find(".author").Text()
		sel.Find(".tags").ChildrenFiltered(".tag").Each(func(i int, tag *goquery.Selection) {
			row.Tags = append(row.Tags, tag.Text())
		})
		rows = append(rows, *row)
	})

	bts, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bts))
}
