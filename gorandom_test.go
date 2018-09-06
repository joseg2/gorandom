package gorandom

import (
	"testing"
)

const succeeded = "✅"
const failed = "❎"

func TestRandom12digits(t *testing.T) {
	t.Log("Given the need to test a random 12 digit number")
	{
		// this just tests got is an int with type assertion
		t.Logf(" Test 0:\t When checking for integer value")
		{
			var (
				got interface{}
				err error
			)

			got, err = Random12digits()
			if err != nil {
				t.Fatalf("\t%s\t Should be able to output an integer: %v", failed, err)
			}

			mob, ok := got.(int)
			if ok {
				t.Logf("\t%s\t Value %d should be an integer: %v", succeeded, mob, ok)
			} else {
				t.Errorf("\t%s\t Value %d should be an integer: %v", failed, mob, ok)
			}
		}
		// TODO implement a better test for random numbers
		t.Logf(" Test 1:\t When checking for short-term non-idempotency")
		{
			const numberOfCycles = 3
			var circularBuffer [numberOfCycles]int

			checkme := 0

			for n, _ := range circularBuffer {
				circularBuffer[n], _ = Random12digits()
				if checkme == circularBuffer[n] {
					t.Errorf("\t%s\t Value returned should be different from previous", failed)
				}
				checkme = circularBuffer[n]
				//t.Logf("\t\t %d-th value: %d",n+1,circularBuffer[n])
			}
			t.Logf("\t%s\t Value returned should be different from previous", succeeded)
		}
	}
}

func TestRandomLine(t *testing.T) {
	t.Log("Given the need to test random lines")
	{
		t.Logf(" Test 1:\t When testing for a random line")
		_, err := RandomLine("/usr/share/dict/words")
		if err != nil {
			t.Fatalf("\t%s\t Should return no errors: %v", failed, err)
		}
		t.Logf("\t%s\t Should return no errors.", succeeded)
		// TODO write a proper type assertion test
	}
}
