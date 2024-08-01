package util

const (
	USD = "USD"
	HKD = "HKD"
	CNY = "CNY"
)

func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, HKD, CNY:
		return true
	default:
		return false
	}
}
