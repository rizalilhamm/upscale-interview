package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	filename := "data_soal_1.txt"

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Gagal membaca file:", err)
		return
	}

	numbersStr := strings.Split(string(content), "\n")
	var sum int
	var numCount = 0

	for _, numStr := range numbersStr {
		if numStr == "" {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sum += num
		numCount++
	}

	for _, numStr := range numbersStr {
		if numStr == "" {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		sum += num
	}

	fmt.Println("Total angka pada file:", numCount)
	fmt.Println("Jumlah semua angka:", sum)
}
