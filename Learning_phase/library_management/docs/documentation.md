# Library Management System - Concurrent Book Reservation

## Overview
This library management system implements concurrent book reservations using Go's concurrency primitives: Goroutines, Channels, and Mutexes.

## Concurrency Architecture

### 1. Goroutines
- **Reservation Worker**: A dedicated goroutine processes all reservation requests from a channel queue
- **Timer Goroutines**: Each reservation spawns a goroutine with a 5-second timer for auto-cancellation

### 2. Channels
- **reserveChan**: Buffered channel (capacity 100) queues incoming reservation requests
- **Reply Channel**: Each request has a reply channel for synchronous error responses
- **Done Channel**: Signals timer goroutines to stop when a book is borrowed

### 3. Mutexes (sync.Mutex)
- Protects shared data structures (books, members, reservationTimers)
- Prevents race conditions during concurrent access
- Used in: AddBook, RemoveBook, BorrowBook, ReturnBook, ListAvailableBooks

## Reservation Flow

1. **Request Submission**: `ReserveBook()` creates a ReservationRequest and sends it to reserveChan
2. **Worker Processing**: The reservation worker goroutine receives and processes the request
3. **Validation**: Checks if book exists and is available (with mutex lock)
4. **Reservation**: Updates book status to "Reserved" and stores memberID
5. **Timer Start**: Launches a 5-second timer goroutine for auto-cancellation
6. **Response**: Returns success/error via reply channel

## Auto-Cancellation Mechanism

```go
timer := time.NewTimer(5 * time.Second)
done := make(chan bool)

go func() {
    select {
    case <-timer.C:
        // Timer expired - cancel reservation
        l.cancelReservation(req.BookID)
    case <-done:
        // Book borrowed in time - stop timer
        return
    }
}()
```

When `BorrowBook()` is called for a reserved book, it closes the done channel to stop the timer.

## Race Condition Prevention

- All map operations (books, members, reservationTimers) are protected by mutex
- Channel communication ensures safe message passing between goroutines
- Timer cleanup prevents memory leaks and ensures proper state management

## Book States

- **Available**: Can be reserved or borrowed
- **Reserved**: Held for a specific member for 5 seconds
- **Borrowed**: Currently checked out by a member

## Key Methods

### ReserveBook(bookID, memberID) error
- Queues reservation request via channel
- Returns immediately with success/error
- Non-blocking operation

### BorrowBook(bookID, memberID) error
- Can borrow reserved books (if reserved by same member)
- Can borrow available books directly
- Cancels reservation timer if applicable

### cancelReservation(bookID)
- Called by timer goroutine after 5 seconds
- Resets book status to Available
- Cleans up timer data

## Concurrent Request Handling

The system safely handles multiple simultaneous reservation attempts:
- Only one member can successfully reserve a book
- Others receive "book not available" error
- First-come-first-served via channel queue
- No data races due to mutex protection

## Testing Concurrent Reservations

See `main.go` for simulation of 10 members concurrently attempting to reserve the same book.
