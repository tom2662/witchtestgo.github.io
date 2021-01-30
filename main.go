package main

import (
	"fmt"

	L "./library"
)

func main() {

	var YearOfDeathA, AgeOfDeathA, YearOfDeathB, AgeOfDeathB int
	var err error

	fmt.Print("Enter Year of Death person A: ")
	_, err = fmt.Scan(&YearOfDeathA)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Print("Enter Age of Death person A: ")
	_, err = fmt.Scan(&AgeOfDeathA)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Print("Enter Year of Death person B: ")
	_, err = fmt.Scan(&YearOfDeathB)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Print("Enter Age of Death Person B: ")
	_, err = fmt.Scan(&AgeOfDeathB)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	var w L.Witchboard = L.RawIngredients{YearOfDeathA, AgeOfDeathA, YearOfDeathB, AgeOfDeathB}

	res := w.KillAverage()

	fmt.Println("Average rate of kill:")
	fmt.Printf("%.2f", res)
}
