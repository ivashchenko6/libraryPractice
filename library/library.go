package library

import (
	"errors"
	"fmt"

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
	//TODO: Fix input, ignore \n symbol by using bufio.Scanner
	fmt.Println("Author Name: ")
	fmt.Scan(&newBook.Author)
	fmt.Println("Book`s Title: ")
	fmt.Scan(&newBook.Title)
	fmt.Println("Book`s Year: ")
	fmt.Scan(&newBook.Year)
	newBook.ID = len(library.Books)

	if isInvalidBook(newBook) {
		return newBook, nil
	}

	return book.Book{}, errors.New("Invalid Book")

}

func isInvalidBook(currentBook book.Book) bool {

	if currentBook.Author != "" && currentBook.Year > 0 && currentBook.Title != "" {
		return true
	}

	return false
}
