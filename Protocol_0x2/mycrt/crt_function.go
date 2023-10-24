package Mycrt 

import (
	"strconv" 
	"crypto/rand"
	"math/bits"
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

func GenShare(sec int, shr int, lim int) []int {
	var ps []int 

	LenShr := bits.OnesCount(uint(sec)) / (lim - 1)   
	source := rand.Reader  
	for i := 0; i < shr; i++ {
		prime, _ := rand.Prime(source, LenShr) 
		ps[i] = prime 
	}
	return ps 


}