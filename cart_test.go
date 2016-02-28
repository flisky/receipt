package receipt

import (
	"os"
	"testing"
)

func TestNewCart(t *testing.T) {
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
			1.11,
			"Product 1",
			"kilo",
		},
	}
	for _, values := range tables {
		db.MustExec(`INSERT INTO product(id, price, name, unitname) values(?, ?, ?, ?)`, values...)
	}
	if _, err := NewCart(map[string]float64{"1": 1}); err != nil {
		t.Errorf("expected no error, got %s", err)
	}
}

func TestMain(m *testing.M) {
	PrepareDB(":memory:")
	os.Exit(m.Run())
}
