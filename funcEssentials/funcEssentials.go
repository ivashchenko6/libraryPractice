package funcEssentials

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"practice.com/example/book"
)

func GetFormattedText(text string) string {
	text = strings.TrimSpace(text)
	return text
}

func IsValidBook(currentBook book.Book) bool {
	if currentBook.Year > 0 && currentBook.Title != "" {
		return true
	}
	return false
}

func GetDataFromUser() (book.Book, error) {

	var newBook book.Book
	reader := bufio.NewReader(os.Stdin)

	//////

	fmt.Print("Book`s Title: ")
	text, _ := reader.ReadString('\n')
	formattedText := GetFormattedText(text)
	newBook.Title = formattedText

	fmt.Print("Book`s Genre: ")
	text, _ = reader.ReadString('\n')
	formattedText = GetFormattedText(text)
	newBook.Genre = formattedText
	/////
	fmt.Print("Book`s  Year: ")
	text, _ = reader.ReadString('\n')
	formattedText = GetFormattedText(text)
	year, err := strconv.Atoi(formattedText)
	if err != nil {
		panic(err)
	}
	newBook.Year = year
	id, _ := gonanoid.Generate("abcdef", 6)
	newBook.ID = id
	if IsValidBook(newBook) {
		return newBook, nil
	}
	return book.Book{}, errors.New("Invalid Book")

}

func GetMultiWordsLineFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fortmattedText := GetFormattedText(text)

	return fortmattedText
}
