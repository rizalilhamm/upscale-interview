package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	// Mengirim permintaan GET ke API
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Membaca data respons dari API
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Mengurai data JSON ke dalam bentuk slice dari struct Post
	var posts []Post
	err = json.Unmarshal(data, &posts)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Menampilkan data dalam format JSON
	jsonData, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
