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
	b := &book.Book
	book.ReadData(*b, "book.txt")

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text_slice := strings.Split(text, " ")
		if text_slice[0] == "add" {
			id := &book.Id
			*b, *id = commands.Add(*b, text_slice, *id)
			fmt.Printf("Added data. Code: %s\n", fmt.Sprintf("%04d", *id-1))
		} else if text_slice[0] == "delete" {
			*b = commands.Del(*b, text_slice[1][:4])
			fmt.Printf("Deleted data. Code: %s\n", text_slice[1][:4])
		} else if text_slice[0][:4] == "list" {
			commands.List(*b)
		}
	}
}
