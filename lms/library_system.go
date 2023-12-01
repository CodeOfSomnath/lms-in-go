package lms

import (
	"fmt"
	"strconv"
	"strings"
)

func basicOptions() {
	println("Options are:")
	println("\t1.Show options")
	println("\t2.Add Books")
	println("\t3.Remove Books")
	println("\t4.Borrow Books")
	println("\t5.Submit Books")
	println("\t6.Exit")
}

func TakeNameAndAuthor() (name, author string) {
	fmt.Print("Enter book's title:")
	fmt.Scanln(&name)
	fmt.Print("Enter book's author:")
	fmt.Scanln(&author)
	return
}

func TakeOption() int {
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