package commands

import (
	"fmt"
	"strconv"
)

func Add(book map[string]map[string]string, input []string, id int) (map[string]map[string]string, int) {
	_, err := strconv.Atoi(input[1])
	if err == nil {
		fmt.Println("Name couldn't only contain numbers.")
		return book, id
	}
	_, err = strconv.Atoi(input[2])
	if err != nil {
		fmt.Println("Phone number couldn't contain letter.")
		return book, id
	}

	book[fmt.Sprintf("%04d", id)] = map[string]string{"name": input[1], "phone": input[2], "email": input[3][:len(input[3])-2]}
	id += 1
	return book, id
}

func Del(book map[string]map[string]string, code string) map[string]map[string]string {
	if _, exist := book[code]; exist {
		delete(book, code)
	} else {
		fmt.Printf("Code %s data don't exists in book.\n", code)
	}
	return book
}

func Find(book map[string]map[string]string, input []string) {
	if input[1] == "code" {
		// fmt.Println(book[input[2][:4]])
		fmt.Printf("%s %s %s %s\n", input[2][:4], book[input[2][:4]]["name"], book[input[2][:4]]["phone"], book[input[2][:4]]["email"])
	} else {
		for key, elem := range book {
			if _, exist := book[key][input[1]]; exist {
				if book[key][input[1]] == input[2][:len(input[2])-2] {
					fmt.Printf("%s %s %s %s\n", key, elem["name"], elem["phone"], elem["email"])
				}
			} else {
				fmt.Println(`Invalid command. Please retry with "find [code/name/phone/email]"`)
			}
		}
	}
}

func List(book map[string]map[string]string) {
	for key, elem := range book {
		fmt.Printf("%s %s %s %s\n", key, elem["name"], elem["phone"], elem["email"])
	}
}
