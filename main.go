package main

import "github.com/CodeOfSomnath/lms-in-go/lms"

func main() {
	b1 := lms.NewBook("Hello", "Somnath")
	// b2 := lms.NewBook("Hello", "Somnath")
	println(b1)
	l1 := []int{1,2,3,4,5,6}
	l2 := l1[0:1]
	l2 = append(l2, l1[2:]...)
	for _, v := range l2 {
		print(v)
	}
}
