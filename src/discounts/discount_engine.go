package discounts

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/Joe-Hendley/go-acme-discount-engine/src/models"
	"github.com/Joe-Hendley/go-acme-discount-engine/src/utils"
)

type DiscountEngine struct {
	LoyaltyCard        bool
	Time               time.Time
	twoForOne          string
	twoForOneSlice     []string
	noBBEDiscountSlice []string
}

func NewDiscountEngine() DiscountEngine {
	twoForOne := "Freddo"

	return DiscountEngine{
		Time:               time.Now(),
		twoForOne:          twoForOne,
		twoForOneSlice:     []string{twoForOne},
		noBBEDiscountSlice: []string{"T-Shirt", "Keyboard", "Drill", "Chair"},
	}
}

func (d *DiscountEngine) ApplyDiscounts(items []models.Item) float64 {
	sort.Sort(models.NameSortedItems(items))

	currentItem := ""
	itemCount := 0

	for idx := range items {
		if items[idx].Name != currentItem {
			currentItem = items[idx].Name
			itemCount = 1
		} else {
			itemCount += 1
			if itemCount == 3 && utils.SliceContainsString(d.twoForOneSlice, currentItem) {
				items[idx].Price = 0
				itemCount = 0
			}
		}
	}

	itemTotal := 0.0

	for idx := range items {
		itemTotal += items[idx].Price
		var daysUntilDate int = int(-(utils.RoundToDay(time.Now()).Sub(utils.RoundToDay(items[idx].LocalDate)).Hours() / 24))
		if items[idx].IsPerishable == false {
			if !utils.SliceContainsString(d.noBBEDiscountSlice, currentItem) {
				if daysUntilDate >= 6 && daysUntilDate <= 10 {
					items[idx].Price = items[idx].Price - (items[idx].Price * (5.0 / 100.0))
				} else {
					if daysUntilDate >= 0 && daysUntilDate <= 5 {
						items[idx].Price = items[idx].Price - (items[idx].Price * (10.0 / 100.0))
					} else {

						if daysUntilDate < 0 {
							items[idx].Price = items[idx].Price - (items[idx].Price * (20.0 / 100.0))
						}
					}

				}
			}
		} else {
			if daysUntilDate == 0 {
				now := d.Time
				fmt.Println(now.Hour())
				if now.Hour() >= 0 && now.Hour() < 12 {
					items[idx].Price = items[idx].Price - (items[idx].Price * (5.0 / 100.0))
				} else {
					if now.Hour() >= 12 && now.Hour() < 16 {
						items[idx].Price = items[idx].Price - (items[idx].Price * (10.0 / 100.0))
					} else {

						if now.Hour() >= 16 && now.Hour() < 18 {
							items[idx].Price = items[idx].Price - (items[idx].Price * (15.0 / 100.0))
						} else {

							if now.Hour() >= 18 && !strings.Contains(items[idx].Name, "(Meat)") {
								items[idx].Price = items[idx].Price - (items[idx].Price * (25.0 / 100.0))
							} else {
								if now.Hour() >= 18 && strings.Contains(items[idx].Name, "(Meat)") {
									items[idx].Price = items[idx].Price - (items[idx].Price * (15.0 / 100.0))
								}
							}

						}
					}

				}
			}
		}
	}

	currentItem = ""
	itemCount = 0

	for idx, item := range items {
		if item.Name != currentItem {
			currentItem = item.Name
			itemCount = 1
		} else {
			itemCount += 1
			if itemCount == 10 && !utils.SliceContainsString(d.twoForOneSlice, item.Name) && item.Price >= 5.00 {
				items[idx].Price = items[idx].Price - items[idx].Price*(2.0/100.0)
				items[idx-1].Price = items[idx-1].Price - items[idx-1].Price*(2.0/100.0)
				items[idx-2].Price = items[idx-2].Price - items[idx-2].Price*(2.0/100.0)
				items[idx-3].Price = items[idx-3].Price - items[idx-3].Price*(2.0/100.0)
				items[idx-4].Price = items[idx-4].Price - items[idx-4].Price*(2.0/100.0)
				items[idx-5].Price = items[idx-5].Price - items[idx-5].Price*(2.0/100.0)
				items[idx-6].Price = items[idx-6].Price - items[idx-6].Price*(2.0/100.0)
				items[idx-7].Price = items[idx-7].Price - items[idx-7].Price*(2.0/100.0)
				items[idx-8].Price = items[idx-8].Price - items[idx-8].Price*(2.0/100.0)
				items[idx-9].Price = items[idx-9].Price - items[idx-9].Price*(2.0/100.0)
				itemCount = 0
			}
		}
	}

	finalTotal := 0.0
	for _, item := range items {
		finalTotal += item.Price
	}

	if d.LoyaltyCard == true && itemTotal >= 50.0 {
		finalTotal = finalTotal - (finalTotal * (2.0 / 100.0))
	}

	return math.Round(finalTotal*100.0) / 100.0
}
