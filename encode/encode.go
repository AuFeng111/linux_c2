package encode

import (
	"fmt"
	"strconv"
)

func Encrypt(bytes []byte) []byte {
	key := byte(0x7F)
	// encrypt the byte slice using XOR encryption
	for i := 0; i < len(bytes); i++ {
		bytes[i] ^= key
	}

	return bytes
}

func Decrypt(bytes []byte) []byte {
	key := byte(0x7F)
	// decrypt the byte slice using XOR encryption
	for i := 0; i < len(bytes); i++ {
		bytes[i] ^= key
	}

	return bytes
}
func Enc(src string) string {
	key := []byte{0x61,0x65,0x6f,0x6e,0x64,0x67,0x5f,0x67,0x6f,0x66}
	var result string
	j := 1
	s := ""
	bt := []rune(src)
	for i := 0; i < len(bt); i++ {
		s = strconv.FormatInt(int64(byte(bt[i])^key[j]), 16)
		if len(s) == 1 {
			s = "0" + s
		}
		result = result + (s)

		//result = result + (s)+"M"
		j = (j + 1) % len(key)
	}
	return result
}
func Dec(src string) string {
	key := []byte{0x61,0x65,0x6f,0x6e,0x64,0x67,0x5f,0x67,0x6f,0x66}
	var result string
	var s int64
	j := 1
	fmt.Println(src)
	bt := []rune(src)
	for i := 0; i < len(src)/2; i++ {
		s, _ = strconv.ParseInt(string(bt[i*2:i*2+2]), 16, 0)
		result = result + string(byte(s)^key[j])
		j = ((j + 1) % len(key))
	}
	return result
}
