package domain

type Response struct {
	ID         int64 `json:"id"`
	CustomerID int64 `json:"customer_id"`
	Accepted   bool  `json:"accepted"`
}
