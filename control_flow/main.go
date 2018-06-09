package main

import "fmt"

func main() {
	for i := 5; i > 0; i -= 2 {
		fmt.Println(i)
	}

	if false {
		fmt.Println("if") // not printed
	} else if true {
		fmt.Println("else if") // printed
	} else {
		fmt.Println("else") // not printed
	}

	youWin()
}

func playHand(score int) string {
	if score < 17 {
		return "hit me"
	} else if score > 21 {
		return "bust"
	} else {
		return "stand"
	}
}

func youWin() {
	fmt.Print("You win: ")
	doorNumber := 1
	switch doorNumber {
	case 1:
		fmt.Println("a new car!") // not printed
		fallthrough               // only the selected case runs by default, unless you includde `fallthrough`
	case 2:
		fmt.Println("a llama!") // printed
	default:
		fmt.Println("a goat!") // not printed
	}
}

func shotsDescription(shots int, par int) string {
	shotsOverPar := shots - par
	switch shotsOverPar {
	case 1:
		return "bogey"
	case 0:
		return "par"
	case -1:
		return "birdie"
	case -2:
		return "eagle"
	default:
		return "no description"
	}
}
