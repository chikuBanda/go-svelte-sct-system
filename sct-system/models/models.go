package models

import "time"

type Beneficiary struct {
	ID            int        `json:"id"`
	Name          string     `json:"name" form:"name" binding:"required"`
	AccountNumber string     `json:"account_number" form:"account_number" binding:"required"`
	Transfers     []Transfer `json:"transfers"` // One-To-Many relationship with 'transfers' table
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

type Transfer struct {
	ID                 int                `json:"id"`
	BeneficiaryId      int                `json:"beneficiary_id" form:"beneficiary_id" binding:"required"`
	Amount             float64            `json:"amount" form:"amount" binding:"required"`
	Status             string             `json:"status"`
	TransactionHistory TransactionHistory `json:"transaction_history"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
}

type TransactionHistory struct {
	ID         int       `json:"id"`
	TransferID int       `json:"transfer_id"`
	Timestamp  time.Time `json:"timestamp"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
