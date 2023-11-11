package Mycrt 

import (
	"testing" 
)


func BenchmarkGen(b *testing.B) {
	CheckAtoi("123") 

	CheckAtoi("abc") 

	ReverseArray([]int{1,2,3,4}) 

	EvalAt([]int{1,2,3,4}, 5, prime) 

	extendedGCD(18, 48)
}

