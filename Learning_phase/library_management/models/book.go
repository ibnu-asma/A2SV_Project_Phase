package models



type BookStatus string

const (
	StatusAvailable BookStatus = "Available"
	StatusBorrowed BookStatus = "Borrowed"
	StatusReserved BookStatus = "Reserved"
)
type Book struct {
	ID int
	Title string
	Author string
	Status BookStatus
	ReservedBy int 
}

