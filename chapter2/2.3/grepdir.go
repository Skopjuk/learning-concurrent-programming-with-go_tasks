package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func readContent(file string, matchString string, dir string) {
	openedFile, err := os.Open(fmt.Sprintf("%s/%s", dir, file))
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
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		trimmedString := strings.TrimRight(file.Name(), "\r\n")

		go readContent(trimmedString, matchString, directory)

		time.Sleep(1 * time.Second)
	}
}
