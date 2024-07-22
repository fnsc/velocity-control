package domain

import "time"

type LoadRequest struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	LoadAmount string `json:"load_amount"`
	Time       string `json:"time"`
}

type Request struct {
	ID         int64
	CustomerID int64
	LoadAmount float64
	Time       time.Time
}
