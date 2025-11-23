package services

import ( 
	"time"
	"sync"
	"library_management/concurrency"
	"library_management/models"
	"errors"
	"fmt"	
)


type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReserveBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	AddMember(member models.Member)
	ListBorrowedBooks(memberID int) []models.Book
}


type Library struct {
	books map[int]models.Book
	members map[int]models.Member
	mu sync.Mutex
	reserveChan chan concurrency.ReservationRequest
	wg sync.WaitGroup
	reservationTimers map[int]timerInfo
}

type timerInfo struct {
    Timer *time.Timer
    Done  chan bool
}


func NewLibrary() *Library {
	l:= &Library{
		books: make(map[int]models.Book),
		members: make(map[int]models.Member),
		reserveChan: make(chan concurrency.ReservationRequest, 100),
		reservationTimers: make(map[int]timerInfo),
	}
	l.wg.Add(1)
	go l.reservationWorker()
	return l

}


func (l *Library) reservationWorker() {
	defer l.wg.Done()

	for req := range l.reserveChan {
		err := l.processReservation(req)
		req.Reply <- err
	}
}


func (l *Library) processReservation(req concurrency.ReservationRequest) error {
	l.mu.Lock()
	book, exists := l.books[req.BookID]
	if !exists {
		l.mu.Unlock()
		return errors.New("book not found")
	}
	if book.Status != models.StatusAvailable {
		l.mu.Unlock()
		return errors.New("book not available")
	}

	// Reserve it
	book.Status = models.StatusReserved
	book.ReservedBy = req.MemberID
	l.books[req.BookID] = book
	l.mu.Unlock()

	// Start 5-second timer
	timer := time.NewTimer(5 * time.Second)
	done := make(chan bool)

	go func() {
		select {
		case <-timer.C:
			l.cancelReservation(req.BookID)
		case <-done:
			// Borrowed in time
			return
		}
	}()

	// Store timer so BorrowBook can stop it
	l.mu.Lock()
	l.reservationTimers[req.BookID] = timerInfo{Timer: timer, Done: done}
	l.mu.Unlock()

	return nil
}






func (l *Library) AddBook(book models.Book) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.books[book.ID] = book
}

func (l *Library) AddMember(member models.Member) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.members[member.ID] = member
}

func (l *Library) RemoveBook(bookID int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.books, bookID)
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



func (l *Library) BorrowBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	if book.Status == models.StatusReserved && book.ReservedBy == memberID {
		if info, exists := l.reservationTimers[bookID]; exists {
			close(info.Done)
			info.Timer.Stop()
			delete(l.reservationTimers, bookID)
		}
		book.Status = models.StatusBorrowed
		l.books[bookID] = book

		member := l.members[memberID]
		member.BorrowedBooks = append(member.BorrowedBooks, book)
		l.members[memberID] = member
		return nil
	}

	if book.Status != models.StatusAvailable {
		return errors.New("book not available")
	}

	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = models.StatusBorrowed
	l.books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member
	return nil
}



func (l *Library) ReserveBook(bookID int, memberID int) error {
	reply := make(chan error)
	req := concurrency.ReservationRequest{
		BookID:   bookID,
		MemberID: memberID,
		Reply:    reply,
	}
	l.reserveChan <- req
	return <-reply
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status != models.StatusBorrowed {
		return errors.New("book is not borrowed")
	}

	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

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

	book.Status = models.StatusAvailable
	l.books[bookID] = book
	return nil
}


func (l *Library) ListAvailableBooks() []models.Book {
	l.mu.Lock()
	defer l.mu.Unlock()

	var available []models.Book
	for _, book := range l.books {
		if book.Status == models.StatusAvailable {
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


func (l *Library) cancelReservation(bookID int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, exists := l.books[bookID]
	if !exists || book.Status != models.StatusReserved {
		return
	}

	book.Status = models.StatusAvailable
	book.ReservedBy = -1
	l.books[bookID] = book
	delete(l.reservationTimers, bookID)

	fmt.Printf("Reservation for book %d expired\n", bookID)
}
