package entities

type User struct {
	ID    int64  `json:"id" dynamodbav:"id"`
	Name  string `json:"name" dynamodbav:"name"`
	Loans []Loan `json:"loans" dynamodbav:"loans"`
}
