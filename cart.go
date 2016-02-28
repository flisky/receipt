package receipt

type Item struct {
	product          *Product
	total            float64
	paid             float64
	discount         float64
	discounts        []*Discount
	quantity         float64
	discountQuantity float64
	price            float64
	discountPrice    float64
}

type Cart []Item
