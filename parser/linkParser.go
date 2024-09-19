package parser

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

// Link struct for saving links
type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) {
	// Create slice of Links
	var links []Link

	// Parse html
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	// Init function for recursion
	var reading func(*html.Node)

	// Traverse html tree
	reading = func(n *html.Node) {
		// Look for a tags
		if n.Type == html.ElementNode && n.Data == "a" {
			// Init Link struct to store result
			var link Link
			// Find link in a tag and save to struct
			for _, a := range n.Attr {
				if a.Key == "href" {
					link.Href = a.Val
				}
			}

			// Find text in a tag and save to struct
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					link.Text = c.Data
				}
			}
			// Add struct to slice of Link structs
			links = append(links, link)
		}
		// Recursion
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			reading(c)
		}

	}
	reading(doc)
	// Print out Link structs in slice
	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
