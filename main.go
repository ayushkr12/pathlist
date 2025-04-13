package main

import (
	"bufio"
	"fmt"
	urlparser "net/url"
	"os"
	"slices"
	"strings"
)

func extractPath(url string) string {
	parsedUrl, _ := urlparser.Parse(url)
	return parsedUrl.Path
}

func extractWords(urlPath string) []string {
	urlPath = strings.Trim(urlPath, "/")
	return strings.Split(urlPath, "/")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var extractedWords []string

	for scanner.Scan() {
		for _, extractedWord := range extractWords(
			extractPath(scanner.Text()),
		) {
			if !slices.Contains(extractedWords, extractedWord) {
				extractedWords = append(extractedWords, extractedWord)
			}
		}
	}

	for _, word := range extractedWords {
		fmt.Println(word)
	}
}
