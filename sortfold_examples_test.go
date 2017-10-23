package sortfold_test

import (
	"fmt"

	"github.com/akutz/sortfold"
)

func ExampleStringsAreSorted() {
	data := []string{"A", "b", "c", "D"}
	fmt.Println(sortfold.StringsAreSorted(data))
	// Output: true
}
