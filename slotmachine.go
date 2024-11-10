package main

import "fmt"

func getName() string {
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
func getBet(balance uint) uint {
    var bet uint
    for true {
        fmt.Printf("Enter the bet amount or 0 to quit (Balance ₹%d): ", balance)
        fmt.Scan(&bet)
        if bet > balance {
            fmt.Print("Bet cannot be greater than balance")
        } else {
            break
        }
    }
    return bet
}


func main() {
	// var balance uint = 2000
	balance := uint(2000)
	getName()
	for balance > 0{
		bet := getBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
	}
	fmt.Printf("You left with : ₹%d\n", balance)
	
}