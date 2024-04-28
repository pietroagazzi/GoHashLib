package utils

// NextPrime returns the next prime number after n.
func NextPrime(n int) int {
	n++
	for !isPrime(n) {
		n++
	}
	return n
}

// isPrime returns true if n is a prime number.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}
