package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(input string) bool {
	// Memeriksa apakah kata atau kalimat palindrom
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your text: ")

	scanner.Scan()
	sentence := strings.ToLower(scanner.Text())

	if isPalindrome(sentence) {
		fmt.Println("Palindrom")
	} else{
		fmt.Println("Bukan palindrom")
	}
}