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

	szs = []string{"b", "c", "a"}
	sortfold.Strings(szs)
	assertEl(t, szs, 0, "a")
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
		t.Errorf("szs[%d]=%s != exp: %s", idx, act, exp)
		t.Fail()
	}
}

func TestStringsAreSorted(t *testing.T) {
	if !sortfold.StringsAreSorted([]string{"A", "b", "c"}) {
		t.Error("mixed case alphabet is sorted")
		t.Fail()
	}
	if sortfold.StringsAreSorted([]string{"b", "a", "c"}) {
		t.Error("single case alphabet is not sorted")
		t.Fail()
	}
	if sortfold.StringsAreSorted(
		[]string{"Ardvark", "ARDVARKS", "2ardvark"}) {
		t.Error("ardvarks are not sorted")
		t.Fail()
	}
}

func TestCompareFold(t *testing.T) {
	assertEq(t, -1, "AAA", "ZZZ")
	assertEq(t, -1, "AAA", "zzz")
	assertEq(t, -1, "aaa", "zzz")
	assertEq(t, 1, "ZZZ", "AAA")
	assertEq(t, 1, "ZZZ", "aaa")
	assertEq(t, 0, "ZZZ", "zzz")
	assertEq(t, -1, "A", "Z")
	assertEq(t, -1, "A", "Y")
	assertEq(t, -1, "A", "B")
	assertEq(t, 1, "A", "1")
	assertEq(t, 0, "αβδ", "ΑΒΔ")
	assertEq(t, 0, "abcdefghijk", "abcdefghij\u212A")
	assertEq(t, 0, "abcdefghijK", "abcdefghij\u212A")
	assertEq(t, 1, "abcdefghijkz", "abcdefghij\u212Ay")
	assertEq(t, -1, "abcdefghijKy", "abcdefghij\u212Az")
}

func TestAZ(t *testing.T) {
	assertEq(t, -1, "A", "Z")
	szs := []string{"Z", "A"}
	sortfold.Strings(szs)
	assertEl(t, szs, 0, "A")
	assertEl(t, szs, 1, "Z")
}

func assertEq(t *testing.T, exp int, a, b string) {
	if act := sortfold.CompareFold(a, b); act != exp {
		t.Errorf(`CompareFold("%s", "%s") = %d, exp=%d`, a, b, act, exp)
		t.FailNow()
	}
}
