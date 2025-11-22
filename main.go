package main

import (
	"fmt"
	"strconv"
	"tazkia/matdisk/bangunruang"
	"tazkia/matdisk/deret"
)

func main() {
	fmt.Println("hello world!")

	x := 10
	y := 11

	var result int = bangunruang.HitungLuasPersegiPanjang(x, y)

	fmt.Println("luas persegi panjang adalah : ", strconv.Itoa(result))

	fmt.Println(" ------------------- ")
	var n int = 5
	var resultFaktorial int64 = deret.Factorial(n)

	fmt.Printf("Faktorial %d adalah %d \n", n, resultFaktorial)

	fmt.Println("---------------")

	var isGanjil bool = deret.IsGanjilOrGenap(n)

	var textResult string = "Ganjil"

	if isGanjil {
		textResult = "Ganjil"
	} else {
		textResult = "Genap"
	}

	fmt.Printf("bilangan %d adalah %s", n, textResult)

	var k int = 50

	fmt.Println(" ----------------- ")
	fmt.Printf("Hitung setiap nilai sebanyak diperulangan range k = %d untuk setiap faktorial dan apakah ganjil atau genap bilang ke i tersebut \n", k)
	fmt.Println(" ----------------- ")

	for i := 1; i <= k; i++ {
		fmt.Printf("  pada suku ke i = %d  Faktorial %d adalah %d \n", i, i, deret.Factorial(i))

		var isGanjil bool = deret.IsGanjilOrGenap(i)
		var textResult string

		if isGanjil {
			textResult = "Ganjil"
		} else {
			textResult = "Genap"
		}
		fmt.Printf("  pada suku ke i = %d  bilangan %d adalah %s \n \n", i, i, textResult)

	}

}
