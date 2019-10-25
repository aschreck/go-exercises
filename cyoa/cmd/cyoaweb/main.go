package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"practice/go-exercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "The port on which to start the CYOA web app")
	fileName := flag.String("file", "gopher.json", "The JSON file with the Gopher story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)
	f, err := os.Open(*fileName)

	if err != nil {
		panic(err)
	}
	story, _ := cyoa.JSONStory(f)
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server at %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
