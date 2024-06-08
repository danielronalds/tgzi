package main

import (
	"fmt"

	"github.com/danielronalds/tgzi/tgzilib"
)

func main() {
	files, err := tgzilib.GetFiles(".", false)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
