package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("recovered: ", v)
			}
		}()

		panic("Panic-panic-panic!")
	}()
	time.Sleep(time.Second)
}
