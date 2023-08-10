package main

import (
	"fmt"
	"unicode"
)

func countLettersAndDigits(s string) (int, int) {
	letterCount := 0
	digitCount := 0

	for _, char := range s {
		if unicode.IsLetter(char) {
			letterCount++
		} else if unicode.IsDigit(char) {
			digitCount++
		}
	}

	return letterCount, digitCount
}

func main() {
	inputString := "Halo123Dunia456"
	letterCount, digitCount := countLettersAndDigits(inputString)

	fmt.Printf("Jumlah huruf: %d\n", letterCount)
	fmt.Printf("Jumlah angka: %d\n", digitCount)
}
