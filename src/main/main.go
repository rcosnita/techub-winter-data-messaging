package main

import (
	"fmt"
	"http"
)

func main() {
	client := http.PublicInstances.Client
	result, err := client.Get("https://www.google.ro")
	if err != nil {
		panic(err)
	}

	fmt.Printf(result)
}
