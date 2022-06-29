package models

import "time"

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

