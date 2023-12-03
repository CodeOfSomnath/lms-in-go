package lms

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// This is file for lms saving api.
// when lms exit the stored file is automatically will deleted for that reason this file saving api
// will help to store files for later use. Thank you

const NameSeparator = "#"

func LoadBooks(lib *Library) bool {
	lines, _ := os.ReadFile("./lms.csv")
	r := strings.NewReader(string(lines))
	csvReader := csv.NewReader(r)
	records, err := csvReader.ReadAll()

	if err != nil {
		panic(err)
	}
	books := make([]Book, 0)
	for _, row := range records {
		book := NewBook(row[0], row[1])
		books = append(books, book)
	}
	lib.storedBooks = books
	return false
}

func SaveLms(lib *Library) {
	file, _ := os.Create("./lms.csv")
	defer file.Close()
	

	for _, book := range lib.storedBooks {
		arr := fmt.Sprintf("%s,%s\n", book.Name, book.Author)
		file.WriteString(arr)
	}

}
