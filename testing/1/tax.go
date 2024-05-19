package tax

func CalculateTax(amount float64) float64 {
	if amount == 0 {
		return 0
	}
	if amount < 1000 {
		return amount * 0.05
	}
	return 5.0
}
