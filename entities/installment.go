package entities

type LoanInstallment struct {
	ID     int64   `json:"id" dynamodbav:"id"`
	Number int     `json:"number" dynamodbav:"number"`
	Amount float64 `json:"amount" dynamodbav:"amount"`
}

func NewInstallment(loanID int64, number int, amount float64) *LoanInstallment {
	return &LoanInstallment{
		ID:     int64(loanID*1000 + int64(number)),
		Number: number,
		Amount: amount,
	}
}
