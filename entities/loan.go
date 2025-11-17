package entities

import "time"

type Loan struct {
	ID           int64             `json:"id" dynamodbav:"id"`
	UserID       int64             `json:"user_id" dynamodbav:"user_id"`
	Amount       float64           `json:"amount" dynamodbav:"amount"`
	InterestRate float32           `json:"interest_rate" dynamodbav:"interest_rate"`
	TermMonths   int               `json:"term_months" dynamodbav:"term_months"`
	Installments []LoanInstallment `json:"installments" dynamodbav:"installments"`
}

func NewLoan(request LoanRequest) *Loan {
	var installments []LoanInstallment
	for i := range request.TermMonths {
		installment := NewInstallment(request.UserID, i+1, request.Amount/float64(request.TermMonths))
		installments = append(installments, *installment)
	}

	return &Loan{
		ID:           int64(time.Now().UnixNano()),
		UserID:       request.UserID,
		Amount:       request.Amount,
		InterestRate: request.InterestRate,
		TermMonths:   request.TermMonths,
		Installments: installments,
	}
}
