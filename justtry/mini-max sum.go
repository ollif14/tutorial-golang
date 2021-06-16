package main

import (
	"fmt"
)
func main() {
	// Write your code here
	var jumlah int64
	var jumlah1 int64
	var angka [5]int64
	for i:=0; i<5 ; i++{
		fmt.Scan(&angka[i])
	}

	max := angka[0]
	min := angka[0]
	for i:=0; i<5 ; i++{
		if max < angka[i]{
			max =angka[i]
		}
		if min > angka[i]{
			min = angka[i]
		}
	}


	for i:=0; i<5 ; i++{
		jumlah = jumlah + angka[i]
		jumlah1 = jumlah1 + angka[i]
	}
	fmt.Println(jumlah-max, jumlah1-min)
}