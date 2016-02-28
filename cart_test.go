package receipt

import "testing"

func TestNewCart(t *testing.T) {
	_, err := NewCart(map[string]float64{"no-exist-product-id": 1})
	if err != errorBadProduct {
		t.Errorf("expected missing product error, got %s", err)
	}
}
