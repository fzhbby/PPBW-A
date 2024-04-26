package main

import (
	"fmt"
	"os"
)

func main() {

	var err error
	fileinfo, err := os.Stat("Kaylatifa")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if fileinfo.IsDir() {
		fmt.Println("Kaylatifa adalah sebuah direktori")
	} else {
		fmt.Println("Kaylatifa adalah sebuah file.")
	}
}
