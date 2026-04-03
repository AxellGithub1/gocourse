package main

import "fmt"

func main() {
	//age := 25
	//if age >= 18 {
	//	fmt.Println("You are an adult!")
	//}

	temperature := 25
	if temperature >= 30 {
		fmt.Println("Its hot outside")
	} else {
		fmt.Println("Its cold outside!")
	}

	//score := 90
	//if score >= 90 {
	//	fmt.Println("Grade A")
	//} else if score >= 80 {
	//	fmt.Println("Grade B")
	//} else if score >= 70 {
	//	fmt.Println("Grade C")
	//} else {
	//	fmt.Println("Grade D")
	//}
	//if score >= 90 {
	//	fmt.Println("Grade A")
	//}
	//if score >= 80 {
	//	fmt.Println("Grade B")
	//}
	//if score >= 70 {
	//	fmt.Println("Grade C")
	//} else {
	//	fmt.Println("Grade D")
	//} // this will make more prints

	//var age int
	//fmt.Scan(&age)
	//if age >= 60 {
	//	fmt.Println("Elder")
	//} else if age >= 20 {
	//	fmt.Println("Adult")
	//} else if age >= 13 {
	//	fmt.Println("Teen")
	//} else {
	//	fmt.Println("Kid")
	//}

	fmt.Println("Your age: ")
	var age int
	fmt.Scan(&age)
	fmt.Println("Have card? (1: yes, 0: No): ")
	var card int
	fmt.Scan(&card)
	if age >= 18 && card == 1 {
		fmt.Println("You are accepted!")
	} else if age >= 18 && card == 0 {
		fmt.Println("Card needed!")
	} else {
		fmt.Println("You cant enter!")
	}
}
