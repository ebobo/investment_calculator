package model

// LegoSet represents a Lego set metadata.
type Report struct {
	ID              int32   `json:"id" db:"id"`
	Client          string  `json:"client" db:"client"`
	TotalInterest   float32 `json:"total_interest" db:"total_interest"`
	PeriodicPayment float32 `json:"periodic_payment" db:"periodic_payment"`
	TotalPayment    float32 `json:"total_payment" db:"total_payment"`
}
