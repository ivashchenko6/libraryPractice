package library

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"practice.com/example/funcEssentials"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"practice.com/example/book"
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

func (l *Library) RemoveBook(id int) {

}

func (l Library) PrintAllBooks() {
	fmt.Println(l.Books)
}

func getDataFromUser(library *Library) (book.Book, error) {

	var newBook book.Book
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Author Name: ")
	text, _ := reader.ReadString('\n')
	formattedText := funcEssentials.GetFormattedText(text)
	newBook.Author = formattedText

	//////

	fmt.Println("Book`s Title: ")
	text, _ = reader.ReadString('\n')
	formattedText = funcEssentials.GetFormattedText(text)
	newBook.Title = formattedText

	/////
	fmt.Println("Book`s  Year: ")
	text, _ = reader.ReadString('\n')
	formattedText = funcEssentials.GetFormattedText(text)
	year, err := strconv.Atoi(formattedText)
	if err != nil {
		panic(err)
	}
	newBook.Year = year
	id, _ := gonanoid.New()
	newBook.ID = id
	if funcEssentials.IsValidBook(newBook) {
		return newBook, nil
	}
	return book.Book{}, errors.New("Invalid Book")

}
