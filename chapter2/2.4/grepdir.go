package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func readContent(file string, matchString string, dir string) {
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

var ListOfFiles []string

func listAllFiles(dir string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			ListOfFiles = append(ListOfFiles, dir+"/"+file.Name())

			continue
		}

		listAllFiles(filepath.Join(dir, file.Name()))
	}

	return ListOfFiles
}

func main() {
	matchString := os.Args[1]

	directory := os.Args[2]

	files := listAllFiles(directory)

	for _, file := range files {
		fmt.Println(file)
		trimmedString := strings.TrimRight(file, "\r\n")

		go readContent(trimmedString, matchString, directory)

		time.Sleep(1 * time.Second)
	}
}
