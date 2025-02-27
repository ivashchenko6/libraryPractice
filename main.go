package main

import (
	"fmt"

	"practice.com/example/library"
)

func main() {
	var lavcLibrary library.Library = library.Library{LibraryName: "LAVC"}
	// lavcLibrary.AddNewBook()
	// lavcLibrary.AddNewBook()
	// lavcLibrary.PrintAllBooks()
	// lavcLibrary.RemoveBook()
	// lavcLibrary.PrintAllBooks()
	fmt.Println("Welcome to LAVC library")
}

func printMenu() {
	fmt.Println("1. Show all books")
	fmt.Println("2. Add new Book")
	fmt.Println("3. Find Book by Name: ")
}
