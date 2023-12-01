package lms


type Library struct {
	storedBooks   []Book
	borrowedBooks []Book
}

// for creating new library
func NewLibrary() Library {
	return Library{
		storedBooks:   make([]Book, 0),
		borrowedBooks: make([]Book, 0),
	}
}

// adding books to library
func (l *Library) AddBook(name, author string) {
	book := NewBook(name, author)
	l.storedBooks = append(l.storedBooks, book)
}

func (l *Library) RemoveBook(name, author string) *Book {
	book := NewBook(name, author)
	for _, b := range l.storedBooks {
		if book == b {
			return &book
		}
	}
	return nil
}

func (l *Library) BorrowBook(name, author string) *Book {
	book := NewBook(name, author)
	for i, b := range l.storedBooks {
		if book == b {
			books := l.storedBooks[:i]
			books = append(books, l.storedBooks[i+1:len(l.storedBooks)]...)
			l.storedBooks = books
			l.storedBooks = append(l.storedBooks, book)
			return &book
		}
	}
	return nil
}

func (l *Library) SubmitBook(name, author string) bool {
	book := NewBook(name, author)
	for i, b := range l.borrowedBooks {
		if book == b {
			books := l.borrowedBooks[:i]
			books = append(books, l.borrowedBooks[i+1:len(l.borrowedBooks)]...)
			l.borrowedBooks = books
			l.storedBooks = append(l.storedBooks, book)
			return true
		}
	}
	return false
}