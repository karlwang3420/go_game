package main

import (
<<<<<<< ours
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
=======
	"fmt"
	"go_game/location"
)

func main() {
	fmt.Println("Hello, World!")

	loc := location.Location{Latitude: 37.7749, Longitude: -122.4194}

	fmt.Println(loc.String())
}
>>>>>>> theirs
package main

import (
	"fmt"
	"go_game/location"
)

func main() {
	fmt.Println("Hello, World!")

	loc := location.Location{Latitude: 37.7749, Longitude: -122.4194}

	fmt.Println(loc.String())
}