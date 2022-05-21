package book

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

// Book that records information
var Book map[string]map[string]string = make(map[string]map[string]string)

// Read file and save in map
func ReadInformation(book map[string]map[string]string, filename string) map[string]map[string]string {
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
			val := fileScanner.Text()
			a := strings.Split(bytes.NewBuffer([]byte(val)).String(), " ")
			book[a[0]] = map[string]string{"name": a[1], "phone": a[2], "email": a[3]}
		}
	}

	return book
}
