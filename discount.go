package receipt

const (
	_ = iota
	DiscountPercentage95
	DiscountBuy2Free1
)

var discountType = map[int]DiscountType{
	DiscountPercentage95: DiscountPercentage{"九五折", 0.95, 0},
	DiscountBuy2Free1:    DiscountFree{"买二赠一", 3, 1},
}

// 当前折扣信息缓存
var discounts []*Discount

type Discount struct {
	id         int
	disabled   []int
	productIds map[ProductId]bool
	DiscountType
}

func (d *Discount) satisfied(item *item) bool {
	if d.productIds[item.product.id] && d.DiscountType.satisfied(item) {
		return true
	}
	return false
}
