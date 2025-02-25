package funcEssentials

import (
	"strings"

	"practice.com/example/book"
)

func GetFormattedText(text string) string {
	text = strings.TrimSpace(text)
	return text
}

func IsValidBook(currentBook book.Book) bool {
	if currentBook.Author != "" && currentBook.Year > 0 && currentBook.Title != "" {
		return true
	}
	return false
}
