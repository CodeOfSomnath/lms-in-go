package lms

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/CodeOfSomnath/lms-in-go/webapi"
)

func showOptions() {
	println("Options are:")
	println("\t1.Show options")
	println("\t2.Add Book")
	println("\t3.Remove Book")
	println("\t4.Borrow Book")
	println("\t5.Submit Book")
	println("\t6.Show Books")
	println("\t7.Exit")

	println("Enter only option number. Like for 'Add Book' enter 1")
}

func TakeNameAndAuthor() (name, author string) {
	fmt.Print("Enter book's title:")
	fmt.Scanln(&name)
	fmt.Print("Enter book's author:")
	fmt.Scanln(&author)
	return
}

func takeOption() int {
	var inp string
	fmt.Println("Enter a option:")
	fmt.Scanln(&inp)
	inp = strings.TrimSpace(inp)
	value, err := strconv.Atoi(inp)
	if err != nil {
		return 1
	}
	return value
}

func AddingBook(l *Library) {
	name, author := TakeNameAndAuthor()
	l.AddBook(name, author)
	fmt.Println("Successfully added this book.")
}

func RemoveBook(l *Library) {
	name, author := TakeNameAndAuthor()
	isRemoved := l.RemoveBook(name, author)
	if isRemoved {
		fmt.Println("Book has successfully removed.")
	} else {
		fmt.Println("This book is not found")
	}
}

func BorrowBook(l *Library) {
	name, author := TakeNameAndAuthor()
	isBorrowed := l.BorrowBook(name, author)
	if isBorrowed {
		fmt.Println("This book has sucessfully borrowed.")
	} else {
		fmt.Println("This book not found..")
	}
}

func SubmitBook(l *Library) {
	name, author := TakeNameAndAuthor()
	if l.SubmitBook(name, author) {
		fmt.Println("Book is sucessfully submitted.")
	} else {
		fmt.Println("This book is not borrowed. Please enter 2 for adding this book")
	}
}

func StartLms() {
	lib := NewLibrary()

	web := webapi.NewWebHandler(8090)
	web.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, ".\\static\\index.html")
	})
	web.Mux.HandleFunc("/ab", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, ".\\static\\addbook.html")
	})
	web.Mux.HandleFunc("/rb", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, ".\\static\\removebook.html")
	})
	web.Mux.HandleFunc("/bb", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, ".\\static\\borrowbook.html")
	})
	web.Mux.HandleFunc("/sb", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, ".\\static\\submitbook.html")
	})
	web.Mux.HandleFunc("/st", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, ".\\static\\showtables.html")
	})
	go web.Start()

	LoadBooks(&lib)
	var option int
	fmt.Println("--------- Welcome to lms cli -----------")
	showOptions()
	for {

		option = takeOption()

		SaveLms(&lib)

		if option == 1 {
			showOptions()
		} else if option == 2 {
			AddingBook(&lib)
		} else if option == 3 {
			RemoveBook(&lib)
		} else if option == 4 {
			BorrowBook(&lib)
		} else if option == 5 {
			SubmitBook(&lib)
		} else if option == 6 {
			LoadBooks(&lib)
			fmt.Println(lib)
		} else {
			break
		}
	}
}
