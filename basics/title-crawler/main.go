package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getTitle(content string) string {
	i := strings.Index(content, "<title>")
	j := strings.Index(content, "</title>")
	if i < 0 || j < 0 {
		return "unknown"
	}
	return string(content[i+7 : j])
}

func connectionWorker(in <-chan string, out chan<- string) {
	resp, err := http.Get(<-in)
	if err != nil {
		fmt.Println("[Error]", err)
		return
	}
	defer resp.Body.Close()
	contentInBytes, _ := ioutil.ReadAll(resp.Body)
	out <- string(contentInBytes)
}

func parseWorker(in <-chan string, out chan<- string) {
	content := <-in
	out <- getTitle(content)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please specify at least one url.")
		return
	}

	// Generate the channels
	urlChannel := make(chan string)
	titleChannel := make(chan string)
	resultChannel := make(chan string)

	// Generate the workers.
	for i := 0; i < len(os.Args)-1; i++ {
		go connectionWorker(urlChannel, titleChannel)
		go parseWorker(titleChannel, resultChannel)
	}

	// Initiate the pipeline by sending the urls acquired from the command
	// line arguments.
	for _, url := range os.Args[1:] {
		fmt.Println("Send url", url)
		urlChannel <- url
	}

	for i := 0; i < len(os.Args)-1; i++ {
		fmt.Println("Received title", <-resultChannel)
	}

}
