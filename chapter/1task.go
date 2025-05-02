package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func readContent(file string) {
	openedFile, err := os.Open(file)
	if err != nil {
		log.Print(err)

		return
	}

	b, err := io.ReadAll(openedFile)
	fmt.Print(string(b))
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
	filenames := os.Args[1:]

	for _, file := range filenames {
		trimmedString := strings.TrimRight(file, "\r\n")

		go readContent(trimmedString)

		time.Sleep(1 * time.Second)
	}
}
