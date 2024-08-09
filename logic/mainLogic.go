package logic

import (
	"bufio"
	"fmt"
	"go_game/state"
	"os"
)

func MainLoop(state state.State) {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		checkHealth(state)

		fmt.Print("Enter text: ")
		scanner.Scan()

		input := scanner.Text()

		fmt.Printf("You entered: %s\n", input)

	}
}

func checkHealth(state state.State) {
	if state.Health < 0 {
		fmt.Println("You are dead")
		os.Exit(0)
	}
}
