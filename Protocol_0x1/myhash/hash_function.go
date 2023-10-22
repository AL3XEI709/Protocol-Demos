package Myhash

import (
	"math/big"
	"crypto/rand"
	"strconv"
)

func EzHash(x string) string{
	return x 
}

func GenKeyB() string{
	KeyB := "" 
	for i := 0; i < 50; i++ {
		tmp, _ := rand.Int(rand.Reader, big.NewInt(100))
		msb := tmp.Int64() % 2
		KeyB = KeyB + strconv.Itoa(int(msb))  
	}
	return KeyB
}

func GenHashC(Nce int, RandX int, KeyB string) string{
	cnt := 0
	HashC := "" 
	for ; Nce > 0; Nce /= 2 {
		lsb := Nce % 2  
		if lsb == 0 {
			HashC = HashC + KeyB[cnt:cnt+1] 
		} else {
			bi, _ := strconv.Atoi(KeyB[cnt:cnt+1]) 
			HashC = HashC + strconv.FormatInt(int64(bi ^ RandX), 2) 
		}
		cnt += 1 
	}
	return HashC
}