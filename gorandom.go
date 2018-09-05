//  Random numbers
//

package gorandom

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Print a random 12 digit number
func Random12digits() {
	const MAX = 999999999999

	// This initializes a seed for 'rand' that will be different at every run, but predictable...
	rand.Seed(time.Now().UnixNano())

	// This returns a random integer from 1 to MAX
	fmt.Println("__________________________________")
	xx := rand.Intn(MAX)
	fmt.Println("Random integer: ", xx)
}

// From the numbers in the given pool, pseudo-random pick one and write it to &mynumber
func RandomDigit(mynumber *int, pool int) {
	rand.Seed(time.Now().UnixNano())
	*mynumber = rand.Intn(pool)
}

// Print a random line from the given path
func RandomLine(path string) {
	var lines []string
	var randy int

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	f.Close()

	maxi := len(lines)

	RandomDigit(&randy, maxi)

	fmt.Println(lines[randy])
}
