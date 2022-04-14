package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON file containing the 'Choose Your Own Adventure' story")
	flag.Parse()

	fmt.Printf("Using the story in %s", *filename)
}
