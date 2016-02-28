package receipt

import (
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type ProductId int64

type Product struct {
	Id       ProductId
	Price    float64
	Name     string
	UnitName string
}

func NewProductId(pid string) (ProductId, error) {
	value, err := strconv.ParseUint(pid, 10, 64)
	if err != nil {
		return ProductId(0), err
	}
	return ProductId(value), nil
}

func fetchProducts(productIds []ProductId) (map[ProductId]*Product, []ProductId) {
	query, args, _ := sqlx.In("SELECT * FROM product WHERE id IN (?)", productIds)
	query = db.Rebind(query)
	rows, err := db.Queryx(query, args...)
	if err != nil {
		log.Printf("DB query product fail %s", err)
		return nil, productIds
	}
	products := make(map[ProductId]*Product)
	for rows.Next() {
		product := Product{}
		if err := rows.StructScan(&product); err != nil {
			log.Printf("%s", err)
		}
		products[product.Id] = &product
	}
	if len(productIds) == len(products) {
		return products, nil
	}
	missing := make([]ProductId, 0, len(productIds)-len(products))
	for _, productId := range productIds {
		if _, ok := products[productId]; !ok {
			missing = append(missing, productId)
		}
	}
	return products, missing
}
