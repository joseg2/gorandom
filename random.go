//  Random numbers
//

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random() {
	proverbs := []string{
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Documentation is for users.",
		"Don't PANIC!",
	}
	var n int

	// This initializes a seed for 'rand' that will be different at every run
	rand.Seed(time.Now().UnixNano())

	// This returns a random integer from 1 to len(proverbs)
	n = rand.Intn(len(proverbs))
	fmt.Println("Random proverb: ", proverbs[n], n)
	fmt.Println("__________________________________")
	xx := rand.Intn(999999999999)
	fmt.Println("Random integer: ", xx)
}

func main() {

	random()

}
