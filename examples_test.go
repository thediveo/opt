package opt_test

import (
	"fmt"

	"github.com/thediveo/opt"
)

func Example_if_truth() {
	fmt.Println(opt.If[string](true).Then("it's true").Else("it's fake!"))
	// Output: it's true
}

func Example_if_fake() {
	fmt.Println(opt.If[string](false).Then("it's true").Else("it's fake!"))
	// Output: it's fake!
}
