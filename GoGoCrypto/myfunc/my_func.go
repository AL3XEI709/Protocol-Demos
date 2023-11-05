package Myfunc 

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "bytes"
	"crypto/sha512" 
	"encoding/base64" 
	"fmt"
)

func check(err error) {
    if err != nil {
        panic(err) 
    }
}

func GetRandBytes(n int) []byte {
	res := make([]byte, n) 
	_, e := rand.Read(res) 
	check(e) 
	return res 
}

func Pad(pt []byte) []byte {
	padlen := aes.BlockSize - (len(pt) % aes.BlockSize) 
	padding := bytes.Repeat([]byte{byte(padlen)}, padlen) 
	return append(pt, padding...) 
}

func Unpad(pt []byte) []byte {
	padlen := int(pt[len(pt)-1]) 
	return pt[:len(pt)-padlen] 
}

func AESEnc(pt []byte, key []byte, iv []byte) ([]byte, error) {
	block, e := aes.NewCipher(key) 
	if e != nil {
		return nil, e 
	}
	pt_ := Pad(pt) 
	fmt.Println("(DEBUG) pt_: ", pt_)
	ct := make([]byte, len(pt_)) 
	m := cipher.NewCBCEncrypter(block, iv) 
	m.CryptBlocks(ct, pt_) 
	
	return ct, nil 

}

func AESDec(ct []byte, key []byte, iv []byte) ([]byte, error) {
	block, e := aes.NewCipher(key) 
	if e != nil {
		return nil, e
	}
	m := cipher.NewCBCDecrypter(block, iv) 
	pt_ := make([]byte, len(ct)) 
	m.CryptBlocks(pt_, ct) 
	
	return Unpad(pt_), nil 

}

func Hash(nonce []byte) string{
	h := sha512.New() 
	h.Write(nonce) 
	nonce_hash := base64.StdEncoding.EncodeToString(h.Sum(nil)) 
	return nonce_hash 
}

