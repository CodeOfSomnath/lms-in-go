package lms

import "fmt"

type Book struct {
	Name   string
	Author string
}

//functions of book

// This function create new book
func NewBook(name, author string) Book {
	return Book{
		Name:   name,
		Author: author,
	}
}

// Checks is book is same or not
func IsSameBook(b1, b2 Book) bool {
	if b1.Author == b2.Author && b2.Name == b1.Name {
		return true
	}
	return false
}

func (b Book) String() string {
	return fmt.Sprintf("{Title: %s, Author: %s}", b.Name, b.Author)
}
