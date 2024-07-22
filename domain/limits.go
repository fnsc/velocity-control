package domain

const (
	DailyLimit     float64 = 5000.0
	WeeklyLimit    float64 = 20000.0
	DailyLoadLimit int     = 3
)

type Limits struct {
	Daily     float64
	Weekly    float64
	DailyLoad int
}

func NewLimits() *Limits {
	return &Limits{
		Daily:     DailyLimit,
		Weekly:    WeeklyLimit,
		DailyLoad: DailyLoadLimit,
	}
}
