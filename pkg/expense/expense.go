package expense

import "time"

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Date        time.Time `json:"date"`
}

func NewExpense(description string, amount int) *Expense {
	return &Expense{
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}
}
