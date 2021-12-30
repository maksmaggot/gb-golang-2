package main

import (
	"fmt"
	"gb-golang-2/client"
)

func main() {
	err := client.NewHTTPStatusError(1)
	if err != nil {
		fmt.Println("error")
	}
}
