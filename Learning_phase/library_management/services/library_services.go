package services

import ( 
	"library_management/models"
	"errors"
)


type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(memberID int, bookID int) error
	ReturnBook(memberID int, bookID int) error
	ListAvailableBooks() []models.Book
	AddMember(member models.Member)
	ListBorrowedBooks(memberID int) []models.Book
}


type Library struct {
	books map[int]models.Book
	members map[int]models.Member
}


func NewLibrary() *Library {
	return &Library{
		books: make(map[int]models.Book),
		members: make(map[int]models.Member),
	}

}


func (l *Library) AddBook(book models.Book) {
	l.books[book.ID] = book
}

func (l *Library) AddMember(member models.Member) {
	l.members[member.ID] = member
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.books, bookID)
	// Also remove from any member's borrowed list
	for mid, member := range l.members {
		updated := []models.Book{}
		for _, b := range member.BorrowedBooks {
			if b.ID != bookID {
				updated = append(updated, b)
			}
		}
		member.BorrowedBooks = updated
		l.members[mid] = member
	}
}



// BorrowBook lets a member borrow a book
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	// Update book status
	book.Status = "Borrowed"
	l.books[bookID] = book

	// Add to member's borrowed books
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member

	return nil
}


// ReturnBook returns a borrowed book
func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book is not borrowed")
	}

	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	// Remove book from member's borrowed list
	updated := []models.Book{}
	found := false
	for _, b := range member.BorrowedBooks {
		if b.ID == bookID {
			found = true
			continue
		}
		updated = append(updated, b)
	}
	if !found {
		return errors.New("member did not borrow this book")
	}

	member.BorrowedBooks = updated
	l.members[memberID] = member

	// Update book status
	book.Status = "Available"
	l.books[bookID] = book

	return nil
}


// ListAvailableBooks returns all available books
func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

// ListBorrowedBooks returns books borrowed by a member
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}

