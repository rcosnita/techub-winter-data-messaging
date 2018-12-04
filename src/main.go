package main

import (
	"fmt"
	"http"
)

func main() {
	client := http.NewClient()
	result, err := client.Get("https://www.google.ro")
	if err != nil {
		panic(err)
	}

	fmt.Printf(result)
}
