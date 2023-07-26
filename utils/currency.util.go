package utils

const (
	USD = "USD"
	EUR = "EUR"
	IDR = "IDR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, IDR:
		return true
	}
	return false
}
