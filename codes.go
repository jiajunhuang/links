package main

import (
	"math"
)

var (
	// EncodeMap lower cases + upper cases + numbers
	EncodeMap = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// MapLength length of map
	MapLength = uint(len(EncodeMap))
	// DecodeMap map char to uint
	DecodeMap = map[byte]uint{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'a': 10,
		'b': 11, 'c': 12, 'd': 13, 'e': 14, 'f': 15, 'g': 16, 'h': 17, 'i': 18, 'j': 19, 'k': 20,
		'l': 21, 'm': 22, 'n': 23, 'o': 24, 'p': 25, 'q': 26, 'r': 27, 's': 28, 't': 29, 'u': 30,
		'v': 31, 'w': 32, 'x': 33, 'y': 34, 'z': 35, 'A': 36, 'B': 37, 'C': 38, 'D': 39, 'E': 40,
		'F': 41, 'G': 42, 'H': 43, 'I': 44, 'J': 45, 'K': 46, 'L': 47, 'M': 48, 'N': 49, 'O': 50,
		'P': 51, 'Q': 52, 'R': 53, 'S': 54, 'T': 55, 'U': 56, 'V': 57, 'W': 58, 'X': 59, 'Y': 60,
		'Z': 61,
	}
)

// Encode encode number to string
func Encode(e uint) string {
	if e == 0 {
		return string(EncodeMap[0])
	}

	result := []byte{}
	for e > 0 {
		remainder := e % MapLength
		e = e / 62
		result = append(result, EncodeMap[remainder])
	}

	return string(result)
}

// Decode decode string to number
func Decode(d string) uint {
	var num uint

	for i := 0; i < len(d); i++ {
		num += uint(math.Pow(float64(MapLength), float64(i))) * DecodeMap[d[i]]
	}

	return num
}
