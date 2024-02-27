package main

import (
	"fmt"

)

const pi = 3.14159

func main() {
	var omlooptijd int
	var afstand int
	var planeet string

	fmt.Println("Welke planeet wil je berekenen")
	fmt.Scan(&planeet)
	fmt.Println("Wat is de afstand van een rondje?")
	fmt.Scan(&afstand)
	fmt.Println("Wat is de omlooptijd in uren?")
	fmt.Scan(&omlooptijd)

	snelheid := afstand / omlooptijd

	fmt.Println("De snelheid van de planeet" +planeet+ "is", snelheid , "km/u")

}

