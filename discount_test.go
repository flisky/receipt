package receipt

import "testing"

func TestDiscount(t *testing.T) {
	PrepareDB(":memory:")
	defer db.Close()

	for _, values := range [][]interface{}{
		{
			ProductId(1),
			1.11,
			"Product 1",
			"kilo",
		},
	} {
		db.MustExec(`INSERT INTO product(id, price, name, unitname) values(?, ?, ?, ?)`, values...)
	}

	for _, values := range [][]interface{}{
		{
			1,
			1,
			"",
			"[1]",
		},
	} {
		db.MustExec(`INSERT INTO discount(id, discounttype, disabled, productids) values(?, ?, ?, ?)`, values...)
	}

	LoadDiscounts()
	if len(Discounts) != 1 {
		t.Errorf("LoadDiscounts fail")
	}
}
