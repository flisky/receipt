package receipt

import "math"

type DiscountType interface {
	satisfied(*item) bool
	process(*item)
}

type PercentageDiscount struct {
	description string
	percentage  float64
	threshold   float64
}

type DiscountFree struct {
	description string
	total       float64
	free        float64
}

func (p *DiscountPercentage) satisfied(item *item) bool {
	if item.paid >= item.threshold {
		return true
	}
	return false
}

func (p *DiscountPercentage) process(item *item) {
	item.discountPrice = item.total * p.percentage
	item.discount += item.discountPrice
	item.paid -= item.discount
}

func (p *DiscountFree) satisfied(item *item) bool {
	if item.quantity >= item.total {
		return true
	}
	return false
}

func (p *DiscountFree) process(item *item) {
	item.discountQuantity = math.Floor(item.quantity * p.free / p.total)
	item.discount += item.discountQuantity * item.price
	item.paid -= item.discountQuantity * item.price
}
