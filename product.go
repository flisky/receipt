package receipt

type ProductId uint64

type Product struct {
	id       ProductId
	price    float64
	Name     string
	UnitName string
}

func fetchProducts(productIds []string) (products map[string]*Product, missing []string) {
	return
}
