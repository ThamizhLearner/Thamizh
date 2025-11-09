package main

import (
	"fmt"

	"github.com/ThamizhLearner/Thamizh/unicode"
)

func main() {
	testUnicode()
}

// Test round-trip: Thamizh Unicode string <=> (Raw) Thamizh letter index slice
func testUnicode() {
	fmt.Println("# Annotated code points dump")
	unicode.DumpAnnotations("தமிழ்")

	fmt.Println()
	fmt.Println("# Thamizh letter index slice")
	fmt.Printf("%v", unicode.Decode("தமிழ்"))

	fmt.Println()
	fmt.Println("# Unicode => Thamizh letter slice => Unicode")
	fmt.Println(unicode.Encode(unicode.Decode("தமிழ்")))
}
