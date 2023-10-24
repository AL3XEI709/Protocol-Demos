package Mycrt 

import "testing" 

func BenchmarkGen(b *testing.B) {
	CheckAtoi("123") 
	CheckAtoi("abc") 
}

