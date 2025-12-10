package entities

type UserRequest struct {
	Name string `json:"name" dynamodbav:"name"`
}
