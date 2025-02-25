package library

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"practice.com/lib/book"
)

type Library struct {
	LibraryName string
	Books       []book.Book
}

func (l *Library) AddNewBook() {
	newBook, err := getDataFromUser(l)
	if err != nil {
		panic(err)
	}

	l.Books = append(l.Books, newBook)
}

func (l Library) PrintAllBooks() {
	fmt.Println(l.Books)
}

func getDataFromUser(library *Library) (book.Book, error) {
	var newBook book.Book
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Author Name: ")
	text, _ := reader.ReadString('\n')
	formattedText := getFormattedText(text)
	newBook.Author = formattedText

	//////

	fmt.Println("Book`s Title: ")
	text, _ = reader.ReadString('\n')
	formattedText = getFormattedText(text)
	newBook.Title = formattedText

	/////
	fmt.Println("Book`s  Year: ")
	text, _ = reader.ReadString('\n')
	formattedText = getFormattedText(text)
	year, err := strconv.Atoi(formattedText)
	if err != nil {
		panic(err)
	}
	newBook.Year = year
	fmt.Println(library.Books)
	if isInvalidBook(newBook) {
		return newBook, nil
	}

	return book.Book{}, errors.New("Invalid Book")

}

func getFormattedText(text string) string {
	text = strings.TrimSpace(text)
	return text
}

func isInvalidBook(currentBook book.Book) bool {

	if currentBook.Author != "" && currentBook.Year > 0 && currentBook.Title != "" {
		return true
	}

	return false
}
