package main

import "fmt"

func main() {
	// Write your code here
	var jumlah int64
	var angka int64
	var size int

	fmt.Scan(&size)
	for i:=0; i< size; i++{
		fmt.Scan(&angka)
		jumlah = jumlah + angka
	}
	fmt.Println(jumlah)
}