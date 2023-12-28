package main

import "github.com/MASTER-REPO-117/library/book"

func main() {
	var myBook = book.Book{
		Title:  "Alicia en el pais de las maravillas",
		Author: "Lewis Carroll",
		Pages:  208,
	}

	myBook.PrintInfo()

}
