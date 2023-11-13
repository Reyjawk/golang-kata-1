package main

import (
//	"bufio"
//	"bytes"
	"encoding/csv"
	"fmt"
//	"io"
	"os"
//	"log"
)
//Error checking for the file
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
	fmt.Println(welcomeMessage())
	books, err := os.Open("/Users/toni/GoLang/golang-kata-1/resources/books.csv")
    check(err)
	defer books.Close()
	/*
	scanFile := bufio.NewScanner(books)
	scanFile.Split(bufio.ScanLines)
	var scanLines []string
	for scanFile.Scan() {
		scanLines = append(scanLines, scanFile.Text())
	}
	for _, line := range scanLines {
		// We only want first and third columns.
		r := NewFieldsReader(line, 0, 1)
		r.Comma = ';'
		recs, err := r.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(recs)
	} */
	//Read the csv file
	reader := csv.NewReader(books)
	//Set delimiter
	reader.Comma = ';'
	//ReadAll in the file
	records, err := reader.ReadAll()
	//Checks for errors
	if err != nil {
		fmt.Println("Error reading records")
	}
	//Loop to iterate through and print each of the string slices
	for _, eachrecord := range records {
		fmt.Println(eachrecord)
	}
}
func welcomeMessage() string {
	return "Hello world!"
}
