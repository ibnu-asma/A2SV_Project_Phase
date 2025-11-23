package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	service services.LibraryManager
	scanner *bufio.Scanner
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{
		service: service,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (c *LibraryController) Start() {
	fmt.Println("Welcome to the Library Management System!")
	for {
		c.showMenu()
		choice := c.readInt("Enter choice")

		switch choice {
		case 1:
			c.addBook()
		case 2:
			c.removeBook()
		case 3:
			c.borrowBook()
		case 4:
			c.returnBook()
		case 5:
			c.listAvailableBooks()
		case 6:
			c.listBorrowedBooks()
		case 7:
			c.addMember()
		case 8:
			c.reserveBook()
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func (c *LibraryController) showMenu() {
	fmt.Println("\n--- Menu ---")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available Books")
	fmt.Println("6. List Borrowed Books")
	fmt.Println("7. Add Member")
	fmt.Println("8. Reserve Book")
	fmt.Println("0. Exit")
}


func (c *LibraryController) readString(prompt string) string {
	fmt.Print(prompt + ": ")
	c.scanner.Scan()
	return strings.TrimSpace(c.scanner.Text())
}

func (c *LibraryController) readInt(prompt string) int {
	for {
		input := c.readString(prompt)
		num, err := strconv.Atoi(input)
		if err == nil {
			return num
		}
		fmt.Println("Please enter a valid number")
	}
}


func (c *LibraryController) addBook() {
	id := c.readInt("Enter Book ID")
	title := c.readString("Enter Title")
	author := c.readString("Enter Author")

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: models.StatusAvailable,
	}
	c.service.AddBook(book)
	fmt.Println("Book added successfully!")
}

func (c *LibraryController) removeBook() {
	id := c.readInt("Enter Book ID to remove")
	c.service.RemoveBook(id)
	fmt.Println("Book removed (if existed)")
}

func (c *LibraryController) borrowBook() {
	bookID := c.readInt("Enter Book ID")
	memberID := c.readInt("Enter Member ID")
	err := c.service.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func (c *LibraryController) returnBook() {
	bookID := c.readInt("Enter Book ID")
	memberID := c.readInt("Enter Member ID")
	err := c.service.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func (c *LibraryController) listAvailableBooks() {
	books := c.service.ListAvailableBooks()
	c.printBooks(books)
}

func (c *LibraryController) listBorrowedBooks() {
	memberID := c.readInt("Enter Member ID")
	books := c.service.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No books borrowed by this member.")
	} else {
		c.printBooks(books)
	}
}

func (c *LibraryController) printBooks(books []models.Book) {
	if len(books) == 0 {
		fmt.Println("No books found.")
		return
	}
	fmt.Printf("%-5s %-20s %-15s %-10s\n", "ID", "Title", "Author", "Status")
	fmt.Println(strings.Repeat("-", 55))
	for _, b := range books {
		fmt.Printf("%-5d %-20s %-15s %-10s\n", b.ID, b.Title, b.Author, b.Status)
	}
}

func (c *LibraryController) addMember() {
	id := c.readInt("Enter Member ID")
	name := c.readString("Enter Member Name")
	member := models.Member{
		ID:            id,
		Name:          name,
		BorrowedBooks: []models.Book{},
	}
	// Type assertion to access AddMember
	if lib, ok := c.service.(*services.Library); ok {
		lib.AddMember(member)
		fmt.Println("Member added successfully!")
	}
}


func (c *LibraryController) reserveBook() {
    bookID := c.readInt("Enter Book ID to reserve")
    memberID := c.readInt("Enter Member ID")

    err := c.service.ReserveBook(bookID, memberID)
    if err != nil {
        fmt.Printf("Reserve failed: %v\n", err)
    } else {
        fmt.Println("Book reserved! Borrow within 5 seconds.")
    }
}