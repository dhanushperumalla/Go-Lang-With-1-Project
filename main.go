package main

import "fmt"
var name string = "Dhanush"
var score int = 50
var names []string

func demo() {
	name := "Dhanush"
	fmt.Print("Hello, world! ")
	fmt.Println("I'm Super Exited to learn Go-Lang")
	fmt.Printf("And I'm %s",name)
	symbols := map[string]uint{
		"A":5,
		"B":10,
		"C":15,
		"D":20,
	}
	multipliers := map[string]uint{
		"A":20,
		"B":15,
		"C":10,
		"D":5,
	}
	fmt.Print("\nA Symbols: ",symbols["A"],"\nA Multipliers: ",multipliers["A"])
}
