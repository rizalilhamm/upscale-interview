package main

import "fmt"

func main() {
    for num := 1; num <= 100; num++ {
        output := ""
        
        if num % 3 == 0 {
            output += "Fizz"
        }
        
        if num % 5 == 0 {
            output += "Buzz"
        }
        
        if output == "" {
            output = fmt.Sprintf("%d", num)
        }
        
        fmt.Println(output)
    }
}
