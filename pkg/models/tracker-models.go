package model

import "time"

// User struct
type User struct {
	ID     int
	Status bool
}

// UserDetail struct
type UserDetail struct {
	ID, Status                int
	Name, Email, Job, Profile string
	Password                  []byte
	CreatedAt, UpdatedAt      time.Time
	DeletedAt                 any
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
	ID, UserID, CateID, Amount   int
	Title, Description, CateName string
	Date, CreatedAt, UpdatedAt   time.Time
}

// Category struct
type Category struct {
	ID          int
	BID         int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   any
}

// CateTotalAmount struct
type CateTotalAmount struct {
	Category                  string
	TotalAmount, BudgetAmount int
}

// Budget struct
type Budget struct {
	ID, UserID int
	Title      string
	Category   string
	Amount     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  any
}

type ExpenseMediaData struct {
	ID, UserID, CateID, Amount int
	Title, Description, ImgURL string
	Date, CreateAt, UpdatedAt  time.Time
}

type MediaData struct {
	Expenses []Expenses
	Media    []ExpenseMediaData
}

type TotalAmountData struct {
	Incomes, Expenses int
}
type DailyAmount struct {
	Day    string
	Amount int
}

type BIData struct {
	Budget   *[]Budget
	ImageSrc string
}
