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

type Cart struct {
	items []*item
}

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
	products, missing := fetchProducts(productIds)
	if missing != nil {
		log.Printf("[Product]cannot find product %v", missing)
		return nil, errorBadProduct
	}
	cart := Cart{}
	cart.items = make([]*item, len(productMap))
	for productId, quantity := range productMap {
		item, err := NewItem(products[productId], quantity)
		if err != nil {
			log.Printf("[Item]init error %s with pid %s & quantity %f", err, productId, quantity)
			return nil, err
		}
		cart.items = append(cart.items, item)
	}
	return &cart, nil
}

func (cart *Cart) Checkout() {
	for _, item := range cart.items {
		// 挑选潜在的折扣
		discounts := make(map[int]*Discount)
		for id, discount := range Discounts {
			if discount.satisfied(item) {
				discounts[id] = discount
			}
		}
		// Enhanced: 健康检查：折扣间的排它是否有冲突
		for id := range discounts {
			if checkConflict(discounts, []int{}, id) {
				ids := make([]int, len(discounts))
				for _, id := range ids {
					ids = append(ids, id)
				}
				log.Printf("[Discount]conflict from %d with discounts %v", id, ids)
			}
		}
		// 剔除不可用的折扣
		for _, discount := range discounts {
			for _, did := range discount.disabled {
				delete(discounts, did)
			}
		}
		// 应用折扣
		for _, discount := range discounts {
			discount.process(item)
		}
	}
}
