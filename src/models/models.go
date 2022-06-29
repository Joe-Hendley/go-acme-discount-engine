// package models is from are ACME internal POS and cannot be modified
package models

import (
	"time"
	"unicode/utf8"
)

type Item struct {
	Name         string
	Price        float64
	IsPerishable bool
	LocalDate    time.Time
}

func NewItem(name string, price float64, isPerishable bool, localDate time.Time) Item {
	return Item{
		Name:         name,
		Price:        price,
		IsPerishable: isPerishable,
		LocalDate:    localDate,
	}
}

type NameSortedItems []Item

func (a NameSortedItems) Len() int      { return len(a) }
func (a NameSortedItems) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a NameSortedItems) Less(i, j int) bool {
	iRune, _ := utf8.DecodeRuneInString(a[i].Name)
	jRune, _ := utf8.DecodeRuneInString(a[j].Name)
	return int32(iRune) < int32(jRune)
}
