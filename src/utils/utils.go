package utils

import (
	"strings"
	"testing"
	"time"

	"github.com/Joe-Hendley/go-acme-discount-engine/src/models"
)

func SliceContainsString(xs []string, y string) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}

func SubSliceContaining(items []models.Item, s string) []models.Item {
	subSlice := []models.Item{}

	for _, item := range items {
		if strings.Contains(item.Name, s) {
			subSlice = append(subSlice, item)
		}
	}

	return subSlice
}

func RoundToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func AssertFloatEquals(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("Got[%f] Want[%f]", got, want)
	}
}