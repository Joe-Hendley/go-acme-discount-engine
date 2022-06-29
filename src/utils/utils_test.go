package utils

import (
	"testing"
	"time"

	"github.com/Joe-Hendley/go-acme-discount-engine/src/models"
)

func TestSubSlice(t *testing.T) {
	items := []models.Item{}
	items = append(items, models.NewItem("Steak (Meat)", 10.00, true, time.Now()))
	items = append(items, models.NewItem("Thing", 6.00, true, time.Now()))
	items = append(items, models.NewItem("T-Shirt", 7.00, false, time.Now().Add(10*24*time.Hour)))

	subSlice := SubSliceContaining(items, "Meat")

	AssertFloatEquals(t, subSlice[0].Price, 10.00)
}

func TestSliceContainsString(t *testing.T) {
	xs := []string{"T-Shirt", "Keyboard", "Drill", "Chair"}

	for _, x := range xs {
		if !SliceContainsString(xs, x) {
			t.Errorf("[%s] not found in slice [%x]", x, xs)
		}
	}
}
