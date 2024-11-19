package main

import (
	"fmt"
)

func main() {
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
	symbolArr := GenerateSymbolArray(symbols)
	
	// var balance uint = 2000
	balance := uint(2000)
	GetName()
	for balance > 0{
		bet := GetBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		spin := getSpin(symbolArr,3,3)
		printSpin(spin)
		winningLines := checkWin(spin,multipliers)
		for i,multi := range winningLines{
			win := multi * bet
			balance += win
			if multi > 0{
				fmt.Printf("Won %d₹, (%dx) on Line #%d\n",win,multi,i+1)
			}
		}
	}
	fmt.Printf("You left with : ₹%d\n", balance)
	
}