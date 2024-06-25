package book

import (
	"fmt"
)

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}

type Textbook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title, author string, pages int, editorial, level string) *Textbook {
	return &Textbook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}
func (b *Book) SetTittle(title string) {
	b.Title = title
}

func (b *Book) GetTitle() string {
	return b.Title

}

func (b *Textbook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nNivel: %s\n", b.Title, b.Author, b.Pages, b.editorial, b.level)

}

type TexBook struct {
	Book
	editorial string
	level     string
}

func NewTexBook(title, Author string, pages int, editorial, level string) *TexBook {
	return &TexBook{
		Book:      Book{title, Author, pages},
		editorial: editorial,
		level:     level,
	}
}
func (b *TexBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nNivel: %s\n",
		b.Title, b.Author, b.Pages, b.editorial, b.level)
}
