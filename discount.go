package receipt

const (
	_ = iota
	DiscountPercentage95
	DiscountBuy2Free1
)

var discountType = map[int]DiscountType{
	DiscountPercentage95: &DiscountPercentage{"九五折", 0.95, 0},
	DiscountBuy2Free1:    &DiscountFree{"买二赠一", 3, 1},
}

// 当前折扣信息缓存
var Discounts map[int]*Discount

type Discount struct {
	id         int
	disabled   []int
	productIds map[ProductId]bool
	DiscountType
}

func (d *Discount) satisfied(item *item) bool {
	if d.productIds[item.product.Id] && d.DiscountType.satisfied(item) {
		return true
	}
	return false
}

func checkConflict(discounts map[int]*Discount, dids []int, id int) bool {
	discount, ok := discounts[id]
	if !ok {
		return false
	}
	dids = append(dids, id)
	for _, did := range discount.disabled {
		if _, ok := discounts[did]; ok {
			if dids[0] == did {
				return true
			}
			if Discounts[did].disabled != nil {
				copies := make([]int, len(dids)+1)
				copy(dids, copies)
				copies = append(copies, did)
				return checkConflict(discounts, copies, did)
			} else {
				return false
			}
		}
	}
	return false
}
