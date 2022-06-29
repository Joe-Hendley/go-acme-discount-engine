# Acme Discount Engine

Acme currently apply discounts manually in-store. Moving online, they want to be able to apply these discounts automatically on checkout on the website. As a bonus, they also want to use this for their Electronic Point of Sale system (EPOS) they use to scan items at tills.

You will need to design a Discount Engine that has the following rules:

## Loyalty Card Scheme

Each customer has a loyalty card that is stamped when they spend over £50. On their 10th instance of spending over £50, there is a discount of 5% on their entire basket for the 10th shop. This discount is applied after any other discounts.

## Special Discounts

2-for-1 offers can be applied to any product. If the product is in a 2 for 1 deal, no other discounts are applied, except the Loyalty Card Scheme if the value of all the items after applying the 2 for 1 discount is above the threshold.

## Bulk Discounts

Any single item worth £5.00 or more, when purchased in a quantity of 10 or more, shall have a discount applied of 2% for the total of the items after any other normal discounts are applied, but before any Loyalty Card discounts. Bulk Discounts do not apply to those items in the 2 for 1 special discount.

## Discounts

### Perishable goods

On the date of use Use-by the item will start to be discounted:
* 12am-12pm - 5% reduction
* 12pm-4pm - 10% reduction
* 4pm-6pm - 15% reduction
* 6pm-Midnight - 25% reduction (does not apply to fresh meat products)

### Non-perishable goods

The Best Before date discounts apply:

* 10 days BB date - 5%
* 5 days BB date - 10%
* After BB date - 20%

These discounts do not apply to clothes, electronic items etc

## Rounding Rules

All rounding is to two decimal places.

We may not use bankers rounding. If the price after discount is a fraction of a penny, then it must always be rounded down if it is .5 or below, otherwise it must be rounded up to the nearest whole penny.

### Examples

* £67.50500 = £67.50
* £67.51500 = £67.51
* £67.52500 = £67.52
* £67.50501 = £67.51
* £67.50750 = £67.51

## Notes

All items are identified by their name only. They are currently not grouped together in any meaningful fashion.

## Constraints

* Because the Item class is shared from the EPOS system, it cannot be altered.
* Sadly, because it was designed by Big Consultancy Group Ltd, they have used Doubles instead of BigDecimals.
* The Acme CEO has heard amazing things about Functional Programming, and the discount engine should produce no side-effects.