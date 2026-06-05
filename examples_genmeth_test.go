//go:build go1.27

package opt_test

import (
	"fmt"

	"github.com/thediveo/opt"
)

func Example_if_truth() {
	fmt.Println(opt.If(true).Then("it's true").Else("it's fake!"))
	// Output: it's true
}

func Example_if_fake() {
	fmt.Println(opt.If(false).Then("it's true").Else("it's fake!"))
	// Output: it's fake!
}
