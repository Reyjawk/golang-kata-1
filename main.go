package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"log"
)
//Error checking for the file
func check(e error) {
    if e != nil {
        panic(e)
    }
}
//Defining CSV Reader
type FieldsReader struct {
	*csv.Reader
	fields []int
}
func NewFieldsReader(r io.Reader, fields ...int) *FieldsReader {
	fr := &FieldsReader{
		Reader: csv.NewReader(r),
		fields: fields,
	}

	return fr
}

func (r *FieldsReader) Read() (record []string, err error) {
	rec, err := r.Reader.Read()
	if err != nil {
		return nil, err
	}

	record = make([]string, len(r.fields))
	for i, f := range r.fields {
		record[i] = rec[f]
	}

	return record, nil
}

func (r *FieldsReader) ReadAll() (records [][]string, err error) {
loop:
	for {
		rec, err := r.Read()
		switch err {
		case io.EOF:
			break loop
		case nil:
			records = append(records, rec)
		default:
			return nil, err
		}
	}

	return records, nil
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
		// We only want first and third columns.
		r := NewFieldsReader(line, 0, 1)
		r.Comma = ';'
		recs, err := r.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(recs)
	}
}

func welcomeMessage() string {
	return "Hello world!"
}
