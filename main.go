package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Enter text: ")
		scanner.Scan()

		input := scanner.Text()

		fmt.Printf("You entered: %s\n", input)

	}

}
