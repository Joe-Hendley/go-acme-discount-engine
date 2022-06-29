// Package utils provides utility functions equivalent to functionality
// in the original Java kata, and should not be modified.
package utils

import (
	"strings"
	"testing"
	"time"

	"github.com/Joe-Hendley/go-acme-discount-engine/src/models"
)

// SliceContainsString returns true if a slice of strings xs contains string y
func SliceContainsString(xs []string, y string) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}

// SubSliceContaining returns a subslice of items with name s from a slice of items
func SubSliceContaining(items []models.Item, s string) []models.Item {
	subSlice := []models.Item{}

	for _, item := range items {
		if strings.Contains(item.Name, s) {
			subSlice = append(subSlice, item)
		}
	}

	return subSlice
}

// RoundToDay rounds the current time to the start of the day
func RoundToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// AssertFloatEquals asserts equality on float64s got, want
func AssertFloatEquals(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("Got[%f] Want[%f]", got, want)
	}
}
