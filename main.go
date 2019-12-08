package main

import (
	"fmt"
	"local/switcher/switcher"
)

func main() {
	errSc := switcher.ScannerLines()
	if errSc != nil {
		fmt.Println(errSc)
	} else {
		fmt.Println("Successfully set")
	}
}
