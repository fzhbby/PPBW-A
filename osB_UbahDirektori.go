package main

import (
	"fmt"
	"os"
)

func main() {

	var err error
	//Mengubah izin direktori.
	err = os.Chmod("Kaylatifa", 0152)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Izin 'Kaylatifa' telah diubah menjadu 0152")

}
