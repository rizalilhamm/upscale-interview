package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)              // Ubah semua karakter menjadi lowercase
	s = strings.Map(func(r rune) rune { // Hapus semua karakter non-alphanumeric
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		return -1
	}, s)

	for i := 0; i < len(s)/2; i++ { // Periksa apakah string sama dengan kebalikannya
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Masukkan string: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("String adalah palindrom.")
	} else {
		fmt.Println("String bukan palindrom.")
	}
}
