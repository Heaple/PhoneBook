package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/heaple/phonebook/book"
	"github.com/heaple/phonebook/commands"
)

func main() {
	b := book.Book
	book.ReadData(b, "book.txt")

	// var s string

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text_slice := strings.Split(text, " ")
		if text_slice[0] == "add" {
			id := book.Id
			b, book.Id = commands.Add(b, text_slice, id)
			fmt.Println(b)
		}
	}
}
