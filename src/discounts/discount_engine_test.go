package discounts_test

import (
	"testing"
	"time"

	"github.com/Joe-Hendley/go-acme-discount-engine/src/discounts"
	"github.com/Joe-Hendley/go-acme-discount-engine/src/models"
)

func assertFloatEquals(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("Got[%f] Want[%f]", got, want)
	}
}

func TestBulkDiscounts(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))

	discountEngine := discounts.NewDiscountEngine()
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	assertFloatEquals(t, result, 49)
}

func TestTwoForOneDiscount(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Freddo", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Freddo", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Freddo", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))

	discountEngine := discounts.NewDiscountEngine()
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	assertFloatEquals(t, result, 10)
}

func TestLoyaltyCard(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 20.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Other Thing", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Thing", 20.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))
	items = append(items, models.NewItem("Something", 5.00, false, time.Date(2022, time.May, 1, 0, 0, 0, 0, time.Local)))

	discountEngine := discounts.NewDiscountEngine()
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	assertFloatEquals(t, result, 49)
}

func TestBestBeforeDiscount(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing1", 10.00, false, time.Now().Add(10*24*time.Hour)))
	items = append(items, models.NewItem("Thing2", 10.00, false, time.Now().Add(6*24*time.Hour)))
	items = append(items, models.NewItem("Thing3", 10.00, false, time.Now().Add(5*24*time.Hour)))
	items = append(items, models.NewItem("Thing4", 10.00, false, time.Now().Add(-1*24*time.Hour)))

	discountEngine := discounts.NewDiscountEngine()
	discountEngine.LoyaltyCard = false
	_ = discountEngine.ApplyDiscounts(items)

	assertFloatEquals(t, items[0].Price, 9.50)
	assertFloatEquals(t, items[1].Price, 9.50)
	assertFloatEquals(t, items[2].Price, 9)
	assertFloatEquals(t, items[3].Price, 8)
}

func TestNoBestBeforeDiscount(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("T-Shirt", 10.00, false, time.Now().Add(10*24*time.Hour)))
	items = append(items, models.NewItem("Keyboard", 10.00, false, time.Now().Add(6*24*time.Hour)))
	items = append(items, models.NewItem("Drill", 10.00, false, time.Now().Add(5*24*time.Hour)))
	items = append(items, models.NewItem("Chair", 10.00, false, time.Now().Add(-1*24*time.Hour)))

	discountEngine := discounts.NewDiscountEngine()
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	assertFloatEquals(t, result, 40)
	assertFloatEquals(t, items[0].Price, 10)
	assertFloatEquals(t, items[1].Price, 10)
	assertFloatEquals(t, items[2].Price, 10)
	assertFloatEquals(t, items[3].Price, 10)
}

func TestUseByDiscountFivePercent(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 10.00, true, time.Now()))

	discountEngine := discounts.NewDiscountEngine()
	discountEngine.LocalDateTime = time.Now().Add(3 * time.Hour)
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	assertFloatEquals(t, items[0].Price, 9.50)
	assertFloatEquals(t, result, 9.50)
}

func TestUseByDiscountTenPercent(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 10.00, true, time.Now()))

	discountEngine := discounts.NewDiscountEngine()
	discountEngine.LocalDateTime = time.Now().Add(13 * time.Hour)
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	assertFloatEquals(t, items[0].Price, 9.0)
	assertFloatEquals(t, result, 9.0)
}

func TestUseByDiscountFifteenPercent(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 10.00, true, time.Now()))

	discountEngine := discounts.NewDiscountEngine()
	discountEngine.LocalDateTime = time.Now().Add(17 * time.Hour)
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	assertFloatEquals(t, items[0].Price, 8.5)
	assertFloatEquals(t, result, 8.5)
}
func TestUseByDiscountTwentyFivePercent(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Thing", 10.00, true, time.Now()))

	discountEngine := discounts.NewDiscountEngine()
	discountEngine.LocalDateTime = time.Now().Add(19 * time.Hour)
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	assertFloatEquals(t, items[0].Price, 7.5)
	assertFloatEquals(t, result, 7.5)
}
func TestUseByDiscountMeatNotTwentyFivePercent(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Steak (Meat)", 10.00, true, time.Now()))

	discountEngine := discounts.NewDiscountEngine()
	discountEngine.LocalDateTime = time.Now().Add(19 * time.Hour)
	discountEngine.LoyaltyCard = false
	result := discountEngine.ApplyDiscounts(items)

	assertFloatEquals(t, items[0].Price, 8.5)
	assertFloatEquals(t, result, 8.5)
}

func TestComplexBasketWithLoyalty(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Steak (Meat)", 10.00, true, time.Now()))
	items = append(items, models.NewItem("Thing", 10.00, true, time.Now()))
	items = append(items, models.NewItem("T-Shirt", 10.00, false, time.Now().Add(10*24*time.Hour)))
	items = append(items, models.NewItem("Thing4", 10.00, false, time.Now().Add(10*24*time.Hour)))
}
