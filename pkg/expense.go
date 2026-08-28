package pkg

import (
	"fmt"
	"time"
)

type Expense struct {
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Date        time.Time `json:"date"`
}

func NewExpense(description string, amount int, timestamp *time.Time) *Expense {
	fmt.Println(description, amount, timestamp)
	if timestamp.IsZero() {
		now := time.Now()
		timestamp = &now
	}
	return &Expense{Description: description, Amount: amount, Date: *timestamp}
}
