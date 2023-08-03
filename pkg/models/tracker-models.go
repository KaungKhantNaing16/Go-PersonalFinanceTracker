package model

import "time"

// User struct
type User struct {
	ID     int
	Status bool
}

// UserDetail struct
type UserDetail struct {
	ID        int
	UserID    int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Income struct
type Income struct {
	ID, UserID, Amount   int
	Title                string
	Description          string
	FileURL              string
	CreatedAt, UpdatedAt time.Time
}

// Expenses struct
type Expenses struct {
	ID, UserID, CateID, Amount int
	Title, Description         string
	Date, CreatedAt, UpdatedAt time.Time
}

// Category struct
type Category struct {
	ID          int
	Status      bool
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   any
}

// Loan struct
type Loan struct {
	ID          int
	Title       int
	Description string
	Amount      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Budget struct
type Budget struct {
	ID          int
	Title       int
	Description string
	Amount      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
