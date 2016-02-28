package receipt

import "errors"

type item struct {
	product  *Product
	quantity float64

	total            float64
	paid             float64
	discount         float64
	discounts        []*Discount
	discountQuantity float64
	price            float64
	discountPrice    float64
}

type Cart []*item

var errorMissing = errors.New("Product not found")

func NewCart(productMap map[string]float64) (*Cart, error) {
	productIds := make([]string, len(productMap))
	for productId := range productMap {
		productIds = append(productIds, productId)
	}
	products, fully := fetchProducts(productIds)
	if fully == false {
		return nil, errorMissing
	}
	cart := make(Cart, len(productMap))
	for productId, quantity := range productMap {
		cart = append(cart, &item{product: products[productId], quantity: quantity})
	}
	return &cart, nil
}

func (c *Cart) Checkout() {
}
