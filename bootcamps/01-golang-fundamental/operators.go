package main

import "fmt"

func main() {
	// Aritmatika
	fmt.Println(1 + 5)  // Penambahan
	fmt.Println(2 - 1)  // Pengurangan
	fmt.Println(2 * 8)  // Perkalian
	fmt.Println(4 / 2)  // Pembagian
	fmt.Println(10 % 3) // Sisa pembagian

	// Augmented Assignment
	var a = 10
	a += 10 // a = 10 + 10 = 20
	fmt.Println(a)

	a -= 10 // a = 20 - 10 = 10
	fmt.Println(a)

	a *= 10 // a = 10 * 10 = 100
	fmt.Println(a)

	a /= 10 // a = 100 / 10 = 10
	fmt.Println(a)

	a %= 10 // a = 10 % 10 = 0
	fmt.Println(a)

	// Unary Operator
	a++ // a = 0 + 1 = 1
	fmt.Println(a)

	a-- // a = 1 - 1 = 0
	fmt.Println(a)

	fmt.Println(-a) // -0
	fmt.Println(+a) // 0

	// Operator Perbandingan
	value := 1 + 1
	isEqual := (value == 2)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	isNotEqual := (value != 2)
	fmt.Printf("nilai %d (%t) \n", value, isNotEqual)

	isLess := (value < 3)
	fmt.Printf("nilai %d (%t) \n", value, isLess)

	isLessOrEqual := (value <= 2)
	fmt.Printf("nilai %d (%t) \n", value, isLessOrEqual)

	isGreater := (value > 1)
	fmt.Printf("nilai %d (%t) \n", value, isGreater)

	isGreaterOrEqual := (value >= 2)
	fmt.Printf("nilai %d (%t) \n", value, isGreaterOrEqual)

	// Operator Logika
	var left = false
	var right = true
	leftAndRight := left && right
	fmt.Printf("left && right =(%t) \n", leftAndRight)

	leftOrRight := left || right
	fmt.Printf("left || right =(%t) \n", leftOrRight)

	leftReverse := !left
	fmt.Printf("!left =(%t) \n", leftReverse)
}
