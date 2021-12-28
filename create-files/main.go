package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 0; i <= 10000; i++ {
		err := createFile(fmt.Sprintf("Filename_%d", i))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

func createFile(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}
