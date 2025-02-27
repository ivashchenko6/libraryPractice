package library

import (
	"errors"
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
	for index, book := range l.Books {
		fmt.Printf("%d. Title: %s, Year: %d\n", index, book.Title, book.Year)
	}
}

func (l Library) FindBookByName() book.Book {
	var lookingForName string
	fmt.Print("What book are we looking for: ")
	fmt.Scan(&lookingForName)

	for _, currentBook := range l.Books {
		if currentBook.Title == lookingForName {
			return currentBook
		}
	}
	return book.Book{}
}

func (l *Library) EditBookName() error {
	var numberOfBook int
	fmt.Print("What book you would like to edit(number): ")
	fmt.Scan(&numberOfBook)
	if numberOfBook > len(l.Books) {
		return errors.New("Out of range.  Must be in the range")
	}

	newTitle := funcEssentials.GetMultiWordsLineFromUser()

	if len(newTitle) < 1 {
		return errors.New("Title of book should be greater than zero")
	}

	l.Books[numberOfBook-1].Title = newTitle
	return nil

}
