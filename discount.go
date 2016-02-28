package receipt

type Discount struct {
	id                 int
	description        string
	disabled           []int
	products           []ProductId
	percentagePrice    float64
	thresholdPrice     float64
	percentageQuantity float64
	thresholdQuantity  float64
}
