package receipt

import "testing"

func TestIntegration(t *testing.T) {
	PrepareDB(":memory:")
	defer db.Close()

	if _, err := NewCart(map[string]float64{"illegal-product-id": 1}); err != errorBadProductId {
		t.Errorf("expected errorBadProductId, got %s", err)
	}
	if _, err := NewCart(map[string]float64{"1": 1}); err != errorBadProduct {
		t.Errorf("expected errorBadProduct, got %s", err)
	}

	tables := [][]interface {
	}{
		{
			ProductId(1),
			1,
			"Product 1",
			"kilogram",
		},
		{
			ProductId(2),
			2,
			"Product 2",
			"Piece",
		},
	}
	for _, values := range tables {
		db.MustExec(`INSERT INTO product(id, price, name, unitname) values(?, ?, ?, ?)`, values...)
	}
	cart, err := NewCart(map[string]float64{"1": 1, "2": 7})
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	discountData := [...][]interface{}{
		{
			1,
			DiscountPercentage95,
			"",
			"[1]",
		},
		{
			2,
			DiscountBuy2Free1,
			"",
			"[2]",
		},
	}
	for _, values := range discountData {
		db.MustExec(`INSERT INTO discount(id, discounttype, disabled, productids) values(?, ?, ?, ?)`, values...)
	}

	LoadDiscounts()
	if len(Discounts) != len(discountData) {
		t.Errorf("LoadDiscounts fail")
	}

	cart.Checkout()
	for _, item := range cart.items {
		switch item.product.Id {
		case ProductId(1):
			if item.paid != 1*1*0.95 {
				t.Errorf("discount percentage calculated error")
			}
		case ProductId(2):
			if item.paid != 2*(7-2) {
				t.Errorf("discount free caculated error")
			}
		}
	}
}
