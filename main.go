package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please specify path to plans file!")
	}

	xmlFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	content, _ := io.ReadAll(xmlFile)
	contentString := string(content)

	routeRegex := regexp.MustCompile("<route .*?</route>")
	linkRegex := regexp.MustCompile("link=\".*?\"")
	out := routeRegex.ReplaceAllString(contentString, "")
	out = linkRegex.ReplaceAllString(out, "")

	outFile, err := os.Create(fmt.Sprintf("cleaned-%s", os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	w := bufio.NewWriter(outFile)
	w.WriteString(out)
}
