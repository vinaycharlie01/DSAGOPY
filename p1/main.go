package main

func sieveOfEratosthenes(min int, max int) []int {
	// Initialize the sieve
	sieve := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		sieve[i] = true
	}

	// Mark non-prime numbers as false in the sieve
	for i := 2; i*i <= max; i++ {
		if sieve[i] {
			for j := i * i; j <= max; j += i {
				sieve[j] = false
			}
		}
	}

	// Collect the prime numbers from the sieve
	primes := []int{}
	for i := min; i <= max; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {

}
