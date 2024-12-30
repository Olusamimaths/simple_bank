package util

const (
	NGN = "NGN"
	USD = "USD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case NGN, USD, EUR:
		return true
	}

	return false
}
