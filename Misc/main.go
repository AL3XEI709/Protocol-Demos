package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

// MignotteThreshold represents the parameters of Mignotte's threshold scheme
type MignotteThreshold struct {
	Prime    *big.Int
	Secret   *big.Int
	Threshold int
	Coeffs   []*big.Int
}

// GenerateCoefficients generates random coefficients for the polynomial
func (mt *MignotteThreshold) GenerateCoefficients() {
	mt.Coeffs = make([]*big.Int, mt.Threshold-1)
	for i := 0; i < mt.Threshold-1; i++ {
		// Generate random coefficients
		coeff := big.NewInt(int64(rand.Intn(100)))
		mt.Coeffs[i] = coeff
	}
}

// EvaluatePolynomial evaluates the polynomial at a given x
func (mt *MignotteThreshold) EvaluatePolynomial(x *big.Int) *big.Int {
	result := new(big.Int).Set(mt.Secret)
	temp := new(big.Int)
	xExp := new(big.Int).Set(x)

	for _, coeff := range mt.Coeffs {
		temp.Exp(x, xExp, mt.Prime)
		temp.Mul(temp, coeff)
		result.Add(result, temp)
	}

	result.Mod(result, mt.Prime)
	return result
}

// SplitSecret generates shares for the given number of participants
func (mt *MignotteThreshold) SplitSecret(numParticipants int) map[int]*big.Int {
	shares := make(map[int]*big.Int)
	for i := 1; i <= numParticipants; i++ {
		x := big.NewInt(int64(i))
		y := mt.EvaluatePolynomial(x)
		shares[i] = y
	}
	return shares
}

// ReconstructSecret reconstructs the secret using the specified shares
func (mt *MignotteThreshold) ReconstructSecret(shares map[int]*big.Int) *big.Int {
	result := new(big.Int)
	for x, y := range shares {
		lagrangeCoeff := mt.computeLagrangeCoefficient(x, shares)
		temp := new(big.Int).Mul(y, lagrangeCoeff)
		result.Add(result, temp)
	}

	result.Mod(result, mt.Prime)
	return result
}

func (mt *MignotteThreshold) computeLagrangeCoefficient(x int, shares map[int]*big.Int) *big.Int {
	result := big.NewInt(1)
	for xi, _ := range shares {
		if xi != x {
			num := new(big.Int).Neg(big.NewInt(int64(xi)))
			denom := new(big.Int).Sub(big.NewInt(int64(xi)), big.NewInt(int64(x)))
			denom.ModInverse(denom, mt.Prime)

			factor := new(big.Int).Mul(num, denom)
			result.Mul(result, factor)
			result.Mod(result, mt.Prime) // Corrected to take the modulus at each step
		}
	}
	return result
}
func main() {
	// Set parameters for Mignotte's threshold scheme
	prime, _ := new(big.Int).SetString("101", 10) // Choose a prime number
	secret := big.NewInt(42)
	threshold := 3

	// Create MignotteThreshold instance
	mignotteScheme := &MignotteThreshold{
		Prime:    prime,
		Secret:   secret,
		Threshold: threshold,
	}

	// Generate random coefficients for the polynomial
	mignotteScheme.GenerateCoefficients()

	// Split the secret into shares
	shares := mignotteScheme.SplitSecret(5)

	// Reconstruct the secret using a subset of shares
	reconstructedSecret := mignotteScheme.ReconstructSecret(map[int]*big.Int{1: shares[1], 3: shares[3], 5: shares[5]})

	// Print results
	fmt.Printf("Original Secret: %s\n", secret.String())
	fmt.Printf("Reconstructed Secret: %s\n", reconstructedSecret.String())
}
