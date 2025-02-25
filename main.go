package main

import (
	"practice.com/lib/library"
)

func main() {

	var lavcLibrary library.Library = library.Library{LibraryName: "LAVC"}
	lavcLibrary.AddNewBook()
	lavcLibrary.PrintAllBooks()
}
