package main

import (
	"flag"
	"fmt"
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

	parsedData := parser.Parse(file)

	// Print out Link structs in slice
	for _, link := range parsedData {
		fmt.Printf("%+v\n", link)
	}
}
