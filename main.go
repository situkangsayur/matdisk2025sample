package main

import (
	"fmt"
	"strconv"
	"tazkia/matdisk/bangunruang"
)

func main() {
	fmt.Println("hello world!")

	x := 10
	y := 11

	var result int = bangunruang.HitungLuasPersegiPanjang(x, y)

	fmt.Println("luas persegi panjang adalah : ", strconv.Itoa(result))
}
