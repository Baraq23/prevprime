package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
	fmt.Println(FindPrevPrime(3))
	fmt.Println(FindPrevPrime(6))
	fmt.Println(FindPrevPrime(9))
	fmt.Println(FindPrevPrime(13))
	// fmt.Println(FindPrevPrime(20))
}

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 1
	}
	primeNums := []int{}
	for i:= 2; i <= nb; i++ {
		if isPrime := PrimeChecker(i); isPrime {
			primeNums = append(primeNums, i)
		}
	}
	return primeNums[len(primeNums)-1]
}

//PrimeChecker() recieves an number and returns true if its a prime number and false otherwise.
func PrimeChecker(n int) bool {
	if n == 2 {
		return true
	}
	hlf := n/2
	for i:= hlf; i >= 2; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}
