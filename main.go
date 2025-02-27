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
		fmt.Print("---> ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			library.PrintAllBooks()
		case 2:
			library.AddNewBook()
		case 3:
			foundedBook, err := library.FindBookByName()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("We found your book, here you go", foundedBook)
		case 4:
			err := library.RemoveBook()
			if err != nil {
				fmt.Println(err)
				continue
			}
		case 5:
			err := library.EditBookName()

			if err != nil {
				fmt.Println(err)
				continue
			}
		case 6:
			fmt.Println("Exit...")
			return
		default:
			fmt.Println("Invalid choice, try again")
			continue
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
