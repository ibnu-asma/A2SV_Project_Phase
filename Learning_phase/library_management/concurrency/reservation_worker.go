package concurrency

type ReservationRequest struct {
	BookID int
	MemberID int
	Reply chan <- error
}