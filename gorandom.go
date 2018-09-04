//  Random numbers
//

package gorandom

import (
	"fmt"
	"math/rand"
	"time"
)

func Random12digits() {
	const MAX = 999999999999

	// This initializes a seed for 'rand' that will be different at every run, but predictable...
	rand.Seed(time.Now().UnixNano())

	// This returns a random integer from 1 to MAX
	fmt.Println("__________________________________")
	xx := rand.Intn(MAX)
	fmt.Println("Random integer: ", xx)
}
