package main

import (
	"github.com/heaple/phonebook/book"
)

func main() {
	b := book.Book
	book.ReadInformation(b, "book.txt")
}
