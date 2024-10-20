package main

// Declare and Print an Array
// Declare an array of 5 integers.
// Initialize the array with values {1, 2, 3, 4, 5}.
// Use a for loop to print each element of the array.

import "fmt"

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}

	// for _, value := range arr {
	// 	fmt.Println(value)
	// }

	// arrFruits := [4]string{"apple", "banana", "cherry", "grape"}
	// fmt.Scanln(&arrFruits[1])
	// for i := 0; i < len(arrFruits); i++ {
	// 	fmt.Print(arrFruits[i] + " ")
	// }
	// fmt.Println(" ")
	// var sum int = 0
	// arrSum := [5]int{1, 2, 3, 4, 5}
	// for i := 0; i < len(arrSum); i++ {
	// 	sum = sum + arrSum[i]

	// }
	// fmt.Println(sum)

	// var arrayInput [5]int
	// for i := 0; i < len(arrayInput); i++ {
	// 	fmt.Scanln(&arrayInput[i])

	// }
	// for i := 0; i < len(arrayInput); i++ {
	// 	fmt.Println(arrayInput[i])
	// }
	// var odd int = 0
	// var even int = 0

	// for i := 0; i < len(arrayInput); i++ {
	// 	if arrayInput[i]%2 == 0 {
	// 		even = even + 1
	// 	} else {
	// 		odd = odd + 1
	// 	}
	// }
	// fmt.Println(even)
	// fmt.Println(odd)

	// arrayEven := [5]int{1, 2, 3, 4, 5}
	// var broi int = 0
	// for i := 0; i < len(arrayEven); i++ {
	// 	if arrayEven[i]%2 == 0 {
	// 		broi = broi + 1
	// 	}

	// }
	// fmt.Println(broi)
	// fmt.Println(len(arrayEven) - broi)

	var arrMax [4]int
	for i := 0; i < len(arrMax); i++ {
		fmt.Scanln(&arrMax[i])
	}
	max := arrMax[0]
	for i := 0; i < len(arrMax); i++ {

		if arrMax[i] > max {
			max = arrMax[i]

		}

	}
	fmt.Print("Nai golemiq element na masiva e: ")
	fmt.Println(max)

}
