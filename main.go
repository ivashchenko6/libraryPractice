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
	chooseAction(&lavcLibrary)
}

func chooseAction(library *library.Library) {
	var choice int

	for {
		printMenu()
		fmt.Scan(&choice)
		if choice < 1 && choice > 6 {
			fmt.Println("Invalid choice, try again")
			continue
		}
		if choice == 6 {
			fmt.Println("Exit")
			return
		}

		switch choice {

		case 1:
			library.PrintAllBooks()
		case 2:
			library.AddNewBook()
		case 3:
			library.FindBookByName()
		case 4:
			library.RemoveBook()
		}
	}

}

func printMenu() {
	fmt.Println("1. Show all books")
	fmt.Println("2. Add a new Book")
	fmt.Println("3. Find Book by Name: ")
	fmt.Println("4. Remove the book")
	fmt.Println("5. Edit title of a book")
	fmt.Println("6. Exit")
}
