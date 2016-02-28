package receipt

import (
	"errors"
	"log"
)

type item struct {
	product  *Product
	quantity float64
	price    float64

	total            float64
	paid             float64
	discount         float64
	discounts        []*Discount
	discountQuantity float64
	discountPrice    float64
}

type Cart []*item

var errorBadProduct = errors.New("Product not found")
var errorBadQuantity = errors.New("Quantity illegal")

func NewItem(product *Product, quantity float64) (*item, error) {
	if quantity <= 0 {
		return nil, errorBadQuantity
	}
	it := item{
		product:  product,
		quantity: quantity,
		price:    product.price,
	}
	it.total = it.quantity * it.price
	it.paid = it.total
	return &it, nil
}

func NewCart(productMap map[string]float64) (*Cart, error) {
	productIds := make([]string, len(productMap))
	for productId := range productMap {
		productIds = append(productIds, productId)
	}
	products, fully := fetchProducts(productIds)
	if fully == false {
		return nil, errorBadProduct
	}
	cart := make(Cart, len(productMap))
	for productId, quantity := range productMap {
		item, err := NewItem(products[productId], quantity)
		if err != nil {
			log.Printf("[Item]init error %s with pid %s & quantity %f", err, productId, quantity)
			return nil, err
		}
		cart = append(cart, item)
	}
	return &cart, nil
}

func (c *Cart) Checkout() {
}
