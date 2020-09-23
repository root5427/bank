package types

// Money represents a monetary amount in minimal units (cents, kopeykas, dirams, etc.).
type Money int64

// Category represents processed payment category
type Category string

type Payment struct {
	ID int
	Amount Money
	Category  Category
}