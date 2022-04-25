package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	file, err := createFile("test.txt")

	// попытка работы с кастомными ошибками
	if errors.Is(err, ErrorCantCreate{}) {
		// вызов явной паники
		panic(time.Now().Format("02-01-2006"))
	}

	filecontent, err := ioutil.ReadAll(file)
	if err != nil {
		// вызов явной паники
		panic(time.Now().Format("02-01-2006"))
	}

	// обработка не явной паники
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Safe exit on panic error:", err)
		}
	}()

	// вызов не явной паники
	fmt.Print("there is panic: %s", filecontent[len(filecontent)])

	defer file.Close()

}

func createFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, ErrorCantCreate{fileName, time.Now().Format("2009-11-10 23:00:00")}
	}

	return file, nil
}

type ErrorCantCreate struct {
	FileName      string
	ErrorDateTime string
}

func (e ErrorCantCreate) Error() string {
	return fmt.Sprintf("File '%s' at %s", e.FileName, e.ErrorDateTime)
}
