package main

import "fmt"

func GetName() string {
	name := ""
	fmt.Println("Welcome to Casino")
	fmt.Print("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome  %s,Lets paly!\n", name)
	return name
}
func GetBet(balance uint) uint {
    var bet uint
    for true {
        fmt.Printf("Enter the bet amount or 0 to quit (Balance ₹%d): ", balance)
        fmt.Scan(&bet)
        if bet > balance {
            fmt.Print("Bet cannot be greater than balance\n")
        } else {
            break
        }
    }
    return bet
}