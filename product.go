package receipt

type ProductId uint64

type Product struct {
	id       ProductId
	Name     string
	UnitName string
	unit     float64
	price    float64
}
