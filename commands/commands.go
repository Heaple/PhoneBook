package commands

import "fmt"

func Add(book map[string]map[string]string, input []string, id int) (map[string]map[string]string, int) {
	book[fmt.Sprintf("%04d", id)] = map[string]string{"name": input[1], "phone": input[2], "email": input[3]}
	id += 1
	return book, id
}