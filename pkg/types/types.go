package types

// Money represents a monetary amount in minimal units (cents, kopeykas, dirams, etc.).
type Money int64

// Category represents processed payment category
type Category string

// Status represents payment status
type Status string

// Predefined payment statuses
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
