package luhn

const (
	luhnMultiplier = 10
)

// CalculateLuhn return the check number
func CalculateLuhn(number int64) int64 {
	checkNumber := checksum(number)

	if checkNumber == 0 {
		return 0
	}

	return luhnMultiplier - checkNumber
}

// Valid check number is valid or not based on Luhn algorithm
func Valid(number int64) bool {
	return (number%luhnMultiplier+checksum(number/luhnMultiplier))%luhnMultiplier == 0
}

func checksum(number int64) int64 {
	var luhn int64

	for i := 0; number > 0; i++ {
		cur := number % luhnMultiplier

		if i%2 == 0 { // even
			cur *= 2
			if cur > luhnMultiplier-1 {
				cur = cur%luhnMultiplier + cur/luhnMultiplier
			}
		}

		luhn += cur
		number /= luhnMultiplier
	}

	return luhn % luhnMultiplier
}
