package main

import (
	"practice.com/example/library"
)

func main() {
	var lavcLibrary library.Library = library.Library{LibraryName: "LAVC"}
	lavcLibrary.AddNewBook()
	lavcLibrary.AddNewBook()
	lavcLibrary.PrintAllBooks()
	lavcLibrary.RemoveBook()
	lavcLibrary.PrintAllBooks()
}
