package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
	fmt.Println(welcomeMessage())
	books, err := os.Open("/Users/toni/GoLang/golang-kata-1/resources/books.csv")
    check(err)
	scanFile := bufio.NewScanner(books)
	scanFile.Split(bufio.ScanLines)
	var scanLines []string
	for scanFile.Scan() {
		scanLines = append(scanLines, scanFile.Text())
	}
	books.Close()
	for _, line := range scanLines {
		fmt.Println(line)
	}
}

func welcomeMessage() string {
	return "Hello world!"
}
