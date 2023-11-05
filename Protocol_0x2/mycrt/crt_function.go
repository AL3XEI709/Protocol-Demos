package Mycrt 

import (
	"strconv" 
	"crypto/rand"
	"math/big"
	"fmt"
)

func CheckAtoi(x_err string) (int, string, string) {
	var feedback string 
	var check string 
	x, err := strconv.Atoi(x_err) 
	if err != nil {
		feedback = "Give me an integer." 
		check = "false" 
	} else {
		if x > (1<<33-1) {
			feedback = "Give me something smaller." 
			check = "false" 
		} else {
			feedback = "Submit ok." 
			check = "true" 
		}
	}
	return x, feedback, check 
}

//source := rand.Reader  
//for i := 0; i < shr; i++ {
//prime, _ := rand.Prime(source, LenShr)  

func isPrime(num *big.Int) bool {
	if num.Cmp(big.NewInt(1)) <= 0 {
		return false
	}
	// Check if the number is prime
	return num.ProbablyPrime(20)
}

func generateRandomPrimes(N int) []*big.Int {
	primes := []*big.Int{}
	for len(primes) < N {
		num, err := rand.Prime(rand.Reader, 1024) // Generate a random 1024-bit prime
		if err != nil {
			fmt.Println("Error generating prime:", err)
			continue
		}
		if isPrime(num) {
			primes = append(primes, num)
		}
	}
	return primes
}

func findPrimes(S *big.Int, N, L int) []*big.Int {
	primes := generateRandomPrimes(N)
	for i := 0; i < N-L+1; i++ {
		product := big.NewInt(1)
		for j := i; j < i+L; j++ {
			product.Mul(product, primes[j])
		}
		if product.Cmp(S) == -1 {
			return primes[i : i+L]
		}
	}
	return nil
}