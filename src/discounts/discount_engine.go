package discounts

import (
	"time"

	"github.com/Joe-Hendley/go-acme-discount-engine/src/models"
)

type DiscountEngine struct {
	LoyaltyCard   bool
	LocalDateTime time.Time
}

func NewDiscountEngine() DiscountEngine {
	return DiscountEngine{
		LocalDateTime: time.Now(),
	}
}

func (d *DiscountEngine) ApplyDiscounts(items []models.Item) float64 {
	return 0
}
