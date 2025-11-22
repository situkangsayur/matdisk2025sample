package deret

func Factorial(n int) int64 {

	var totalPerkalian int64 = 1
	// totalPerkalian := 1

	for i := 1; i <= n; i++ {
		totalPerkalian = totalPerkalian * int64(i)
	}

	return totalPerkalian
}

func IsGanjilOrGenap(n int) bool {
	result := true

	if n%2 == 1 {

		result = true

	} else {

		result = false

	}

	return result
}
