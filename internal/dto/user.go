package dto

type UserCreate struct {
	Amount   int64  `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency" binding:"required,oneof=KZT USD"`
}
