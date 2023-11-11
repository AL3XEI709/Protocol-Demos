package Mycrt 

import (
	"strconv" 
	"crypto/rand"
	"math/big"
	//"fmt"
)

var prime int = 4294967291

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

func ReverseArray(a []int) []int {
	l := len(a) 
	rev := make([]int, l) 
	for i := 0; i < l; i ++ {
		rev[i] = a[l-1-i]
	}
	return rev 
}

func EvalAt(poly []int, x int, prime int) int {
	accum := 0 
	poly = ReverseArray(poly) 
	for _, coeff := range poly {
		accum *= x 
		accum += coeff 
		accum %= prime 
	}
	return accum 
}

func extendedGCD(a, b int) (int, int) {
	var quot int 
    x := 0 
	last_x := 1 
	y := 1
	last_y := 0 
	for {
		if b == 0 {
			break
		}
		quot = a/b 
		a,b = b, a%b 
		x, last_x = last_x - quot * x, x
        y, last_y = last_y - quot * y, y
	}
	return last_x, last_y 

}

func MakeRandomShares(secret, minimum, shares int) []int {
	var poly []int 
	var points []int 
	poly[0] = secret 
	for i := 0; i < minimum-1; i ++ {
		poly[i], _ = rand.Int(rand.Reader, big.NewInt(prime-1)) 
	}
}