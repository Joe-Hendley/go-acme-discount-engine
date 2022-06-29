package discounts

import (
	"testing"
	"time"

	"github.com/Joe-Hendley/go-acme-discount-engine/src/models"
	"github.com/Joe-Hendley/go-acme-discount-engine/src/utils"
)

func TestBulkDiscounts(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 5.0, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.0, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.0, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.0, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.0, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.0, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.0, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.0, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.0, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.0, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))

	discountEngine := NewDiscountEngine()
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	utils.AssertFloatEquals(t, result, 49)
}

func TestTwoForOneDiscount(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Freddo", 5.00, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Freddo", 5.00, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Freddo", 5.00, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))

	discountEngine := NewDiscountEngine()
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	utils.AssertFloatEquals(t, result, 10)
}

func TestLoyaltyCard(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 20.00, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Other Thing", 5.00, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 20.00, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Something", 5.00, false, time.Date(2200, time.May, 1, 0, 0, 0, 0, time.Local)))

	discountEngine := NewDiscountEngine()
	discountEngine.LoyaltyCard = true
	result := discountEngine.ApplyDiscounts(items)

	utils.AssertFloatEquals(t, result, 49)
}

func TestBestBeforeDiscount(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing1", 10.00, false, time.Now().Add(10*24*time.Hour)))
	items = append(items, models.NewItem("Thing2", 10.00, false, time.Now().Add(6*24*time.Hour)))
	items = append(items, models.NewItem("Thing3", 10.00, false, time.Now().Add(5*24*time.Hour)))
	items = append(items, models.NewItem("Thing4", 10.00, false, time.Now().Add(-1*24*time.Hour)))

	discountEngine := NewDiscountEngine()
	discountEngine.LoyaltyCard = false
	_ = discountEngine.ApplyDiscounts(items)

	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "Thing1")[0].Price, 9.50)
	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "Thing2")[0].Price, 9.50)
	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "Thing3")[0].Price, 9)
	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "Thing4")[0].Price, 8)
}

func TestNoBestBeforeDiscount(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("T-Shirt", 10.00, false, time.Now().Add(10*24*time.Hour)))
	items = append(items, models.NewItem("Keyboard", 10.00, false, time.Now().Add(6*24*time.Hour)))
	items = append(items, models.NewItem("Drill", 10.00, false, time.Now().Add(5*24*time.Hour)))
	items = append(items, models.NewItem("Chair", 10.00, false, time.Now().Add(-1*24*time.Hour)))

	discountEngine := NewDiscountEngine()
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	utils.AssertFloatEquals(t, result, 40)
	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "T-Shirt")[0].Price, 10)
	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "Keyboard")[0].Price, 10)
	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "Drill")[0].Price, 10)
	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "Chair")[0].Price, 10)
}

func TestUseByDiscountFivePercent(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 10.00, true, time.Now()))

	discountEngine := NewDiscountEngine()
	discountEngine.Time = utils.RoundToDay(time.Now()).Add(3 * time.Hour)
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "Thing")[0].Price, 9.50)
	utils.AssertFloatEquals(t, result, 9.50)
}

func TestUseByDiscountTenPercent(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 10.00, true, time.Now()))

	discountEngine := NewDiscountEngine()
	discountEngine.Time = utils.RoundToDay(time.Now()).Add(13 * time.Hour)
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "Thing")[0].Price, 9.0)
	utils.AssertFloatEquals(t, result, 9.0)
}

func TestUseByDiscountFifteenPercent(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 10.00, true, time.Now()))

	discountEngine := NewDiscountEngine()
	discountEngine.Time = utils.RoundToDay(time.Now()).Add(17 * time.Hour)
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "Thing")[0].Price, 8.5)
	utils.AssertFloatEquals(t, result, 8.5)
}
func TestUseByDiscountTwentyFivePercent(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 10.00, true, time.Now()))

	discountEngine := NewDiscountEngine()
	discountEngine.Time = utils.RoundToDay(time.Now()).Add(19 * time.Hour)
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "Thing")[0].Price, 7.5)
	utils.AssertFloatEquals(t, result, 7.5)
}
func TestUseByDiscountMeatNotTwentyFivePercent(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Steak (Meat)", 10.00, true, time.Now()))

	discountEngine := NewDiscountEngine()
	discountEngine.Time = utils.RoundToDay(time.Now()).Add(19 * time.Hour)
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	utils.AssertFloatEquals(t, utils.SubSliceContaining(items, "(Meat)")[0].Price, 8.5)
	utils.AssertFloatEquals(t, result, 8.5)
}

func TestComplexBasketWithLoyalty(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Steak (Meat)", 10.00, true, time.Now()))
	items = append(items, models.NewItem("Thing", 10.00, true, time.Now()))
	items = append(items, models.NewItem("T-Shirt", 10.00, false, time.Now().Add(10*24*time.Hour)))
	items = append(items, models.NewItem("Thing4", 10.00, false, time.Now().Add(10*24*time.Hour)))
}
