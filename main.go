package main

import (
	"flag"
	"log"
	"os"

	"github.com/demartinom/link/parser"
)

func main() {
	// Test HTML
	htmlToRead := flag.String("doc", "ex1.html", "html to read")
	flag.Parse()

	// Open HTML file
	file, err := os.Open(*htmlToRead)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	parser.Parse(file)
}
