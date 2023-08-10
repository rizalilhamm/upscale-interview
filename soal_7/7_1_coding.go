package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	apiURL := "https://example.com/api/data" // Ganti dengan URL API yang sesuai

	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Gagal melakukan permintaan ke API:", err)
		return
	}
	defer response.Body.Close()

	var data Data
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Gagal memparsing JSON:", err)
		return
	}

	fmt.Println("Nama:", data.Name)
	fmt.Println("Email:", data.Email)
}
