package entities

type LoanRequest struct {
	Amount       float64 `json:"amount"`
	UserID       int64   `json:"user_id"`
	TermMonths   int     `json:"term_months"`
	InterestRate float32 `json:"interest_rate"`
}
