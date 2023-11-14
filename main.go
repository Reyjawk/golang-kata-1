package main

import (
	"bufio"
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
type TitleISBN struct {
	Title string
	ISBN string
}
func createTitleISBNList(records [] []string) []TitleISBN {
	var titleisbnList []TitleISBN
	for i, line := range records {
		if i > 0 { // omit header line
			var rec TitleISBN
			for j, field := range line {
				if j == 0 {
					rec.Title = field
				} else if j == 1 {
					rec.ISBN = field
				}
			}
			titleisbnList = append(titleisbnList, rec)
		}
	}
	return titleisbnList
}
func welcomeMessage() string {
	return "Hello Welcome to Book Lookup Tool!"
}

func main() {
	fmt.Println(welcomeMessage())
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	check(err)
	fmt.Printf("Input was: %s-\n", line)
	books, err := os.Open("/Users/toni/GoLang/golang-kata-1/resources/books.csv")
    check(err)
	defer books.Close()
	fmt.Println("Input T to look up by Title or I to look up by ISBN:")
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
	/*Loop to iterate through and print each of the string slices
	for _, eachrecord := range records {
		fmt.Println(eachrecord)
	} */
	titleisbnList := createTitleISBNList(records)
	//print array
	fmt.Printf("%+v\n", titleisbnList)
}
