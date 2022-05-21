package book

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

// Book that records information
var Book map[string]map[string]string = make(map[string]map[string]string)
var Id int

// Read file and save in map
func ReadData(book map[string]map[string]string, filename string) map[string]map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		os.Create(filename)
		file, _ = os.Open(filename)
	}
	defer file.Close()

	buff := make([]byte, 1024)
	cnt, _ := file.Read(buff)
	file.Seek(0, 0)
	if cnt != 0 {
		fileScanner := bufio.NewScanner(file)

		for fileScanner.Scan() {
			data := fileScanner.Text()
			data_buffer := bytes.NewBuffer([]byte(data))
			if len(data) == 1 {
				Id, _ = strconv.Atoi(data_buffer.String())
			}
			slice := strings.Split(data_buffer.String(), " ")
			book[slice[0]] = map[string]string{"name": slice[1], "phone": slice[2], "email": slice[3]}
		}
	} else {
		Id = 1
	}

	return book
}
