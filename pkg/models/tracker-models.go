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
	ID          int
	UserID      int
	Title       string
	Amount      int
	Description string
	FileURL     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Expenses struct
type Expenses struct {
	ID          int
	UserID      int
	CateID      int
	Title       string
	Amount      int
	Description string
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Category struct
type Category struct {
	ID          int
	Status      bool
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
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
