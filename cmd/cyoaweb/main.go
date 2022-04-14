package main

import (
	"flag"
	"fmt"
	"os"

	cyoa "github.com/Marvellous-Chimaraoke/choose-your-own-adventure"
)

func main() {
	filename := flag.String("file", "gopher.json", "The JSON file containing the 'Choose Your Own Adventure' story")
	flag.Parse()

	fmt.Printf("Using the story in %s...\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

}
