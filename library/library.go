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

func (l *Library) RemoveBook() error {
	var id string
	l.PrintAllBooks()
	fmt.Println("Book need to be deleted(by ID): ")
	fmt.Scan(&id)
	for index, currentBook := range l.Books {
		if currentBook.ID == id {
			l.Books = append(l.Books[:index], l.Books[index+1:]...)
			fmt.Println(currentBook, " Was deleted")
			return nil
		}
	}
	return errors.New("Book wasn`t founded/Invalid book")
}

func (l Library) PrintAllBooks() {
	fmt.Println("--------------------")
	fmt.Println("Current books: ")
	for index, book := range l.Books {
		fmt.Printf("%d. Title: %s, Year: %d, Genre: %s [ID: %s]\n", index+1, book.Title, book.Year, book.Genre, book.ID)
	}
	fmt.Println("--------------------")
}

func (l Library) FindBookByName() (book.Book, error) {
	fmt.Print("What book are we looking for: ")
	var lookingForName string = funcEssentials.GetMultiWordsLineFromUser()

	for _, currentBook := range l.Books {
		if currentBook.Title == lookingForName {
			return currentBook, nil
		}
	}
	return book.Book{}, errors.New("Failure to find the book")
}

func (l *Library) EditBookName() error {
	var numberOfBook int
	fmt.Print("What book you would like to edit(number): ")
	fmt.Scan(&numberOfBook)
	if numberOfBook > len(l.Books) {
		return errors.New("Out of range.  Must be in the range")
	}
	fmt.Print("What is the new title for this book: ")
	newTitle := funcEssentials.GetMultiWordsLineFromUser()

	if len(newTitle) < 1 {
		return errors.New("Title of book should be greater than zero")
	}
	oldTitle := l.Books[numberOfBook-1].Title
	l.Books[numberOfBook-1].Title = newTitle
	fmt.Printf("Title was change from [%s] -------> [%s]\n", oldTitle, newTitle)
	return nil

}
