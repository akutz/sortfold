package sortfold_test

import (
	"testing"

	"github.com/akutz/sortfold"
)

func TestStrings(t *testing.T) {
	szs := []string{"b", "c", "A"}
	sortfold.Strings(szs)
	assertEl(t, szs, 0, "A")
	assertEl(t, szs, 1, "b")
	assertEl(t, szs, 2, "c")

	szs = []string{"Ardvark", "ARDVARKS", "2ardvark"}
	sortfold.Strings(szs)
	assertEl(t, szs, 0, "2ardvark")
	assertEl(t, szs, 1, "Ardvark")
	assertEl(t, szs, 2, "ARDVARKS")
}

func assertEl(t *testing.T, szs []string, idx int, exp string) {
	if act := szs[idx]; act != exp {
		t.Errorf("szs[%d] != exp: %s", act)
		t.Fail()
	}
}

func TestStringsAreSorted(t *testing.T) {
	if !sortfold.StringsAreSorted([]string{"A", "b", "c"}) {
		t.Error("alphabet is sorted")
		t.Fail()
	}
	if sortfold.StringsAreSorted(
		[]string{"Ardvark", "ARDVARKS", "2ardvark"}) {
		t.Error("ardvarks are not sorted")
		t.Fail()
	}
}
