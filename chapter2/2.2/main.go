package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func readContent(file string, matchString string) {
	openedFile, err := os.Open(file)
	if err != nil {
		log.Print(err)

		return
	}

	b, err := io.ReadAll(openedFile)
	fileContent := string(b)
	if strings.Contains(fileContent, matchString) {
		fmt.Printf("%s: found\n", file)

		return
	}
	fmt.Print("\n")

	defer func(openedFile *os.File) {
		err := openedFile.Close()
		if err != nil {
			log.Fatal(err)

		}
	}(openedFile)

	return
}

func main() {
	matchString := os.Args[1]

	directory := os.Args[2]

	files, err := os.ReadDir(directory)
	if err != nil {
		return
	}

	for _, file := range files {
		trimmedString := strings.TrimRight(file.Name(), "\r\n")

		go readContent(trimmedString, matchString)

		time.Sleep(1 * time.Second)
	}
}
