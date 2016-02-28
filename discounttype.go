package receipt

import "math"

type DiscountType interface {
	satisfied(*item) bool
	process(*item)
}

type DiscountPercentage struct {
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
	if item.Paid >= p.threshold {
		return true
	}
	return false
}

func (p *DiscountPercentage) process(item *item) {
	item.DiscountPrice = item.Total * (1 - p.percentage)
	item.Discount += item.DiscountPrice
	item.Paid -= item.Discount
}

func (p *DiscountFree) satisfied(item *item) bool {
	if item.Quantity >= p.total {
		return true
	}
	return false
}

func (p *DiscountFree) process(item *item) {
	item.DiscountQuantity = math.Floor(item.Quantity * p.free / p.total)
	item.Discount += item.DiscountQuantity * item.Product.Price
	item.Paid -= item.DiscountQuantity * item.Product.Price
}
