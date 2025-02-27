package library

import (
	"fmt"

	"practice.com/example/funcEssentials"

	"practice.com/example/book"
)

type Library struct {
	LibraryName string
	Books       []book.Book
}

func (l *Library) AddNewBook() {
	newBook, err := funcEssentials.GetDataFromUser()
	if err != nil {
		panic(err)
	}

	l.Books = append(l.Books, newBook)
}

func (l *Library) RemoveBook() {
	var id string
	fmt.Println("Id need to be deleted: ")
	fmt.Scan(&id)
	for index, currentBook := range l.Books {
		if currentBook.ID == id {
			l.Books = append(l.Books[:index], l.Books[index+1:]...)
			fmt.Println(currentBook, " Was deleted")
		}
	}
}

func (l Library) PrintAllBooks() {
	fmt.Println(l.Books)
}
