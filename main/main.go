// Test station for development of the தமிழ் (Thamizh) package
package main

import (
	"fmt"

	script "github.com/ThamizhLearner/Thamizh"
	"github.com/ThamizhLearner/Thamizh/internal/unicode"
)

func main() {
	testUnicode()
	testLetter()
}

func testLetter() {
	s := script.MustDecode("தமிழ்")
	fmt.Println(s)
	for l := range s.Values() {
		fmt.Println(l)
	}
}

// Test round-trip: Thamizh Unicode string <=> (Raw) Thamizh letter index slice
func testUnicode() {
	fmt.Println("# Annotated code points dump")
	unicode.DumpAnnotations("தமிழ்")

	fmt.Println()
	fmt.Println("# Thamizh letter index slice")
	fmt.Printf("%v\n", unicode.Decode("தமிழ்"))

	fmt.Println()
	fmt.Println("# Thamizh letter slice => Unicode")
	fmt.Println(unicode.Encode(unicode.Decode("தமிழ்")))
}
