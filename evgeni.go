package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sn := rand.Intn(100) + 1
	fmt.Println(sn)
	var input int

	for sn != input {
		fmt.Println("try to guess the number")
		fmt.Scanln(&input)

		if input < sn {
			fmt.Println("too low")
		}
		if input > sn {
			fmt.Println("too high")
		}

	}
	fmt.Println("Ti pozna chisloto!!")

}
