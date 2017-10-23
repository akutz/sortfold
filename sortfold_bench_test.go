package sortfold_test

import (
	"math/rand"
	"sort"
	"strings"
	"testing"

	"github.com/akutz/sortfold"
)

////////////////////////////////////////////////////////////////////////////////
//                                INPUT                                       //
////////////////////////////////////////////////////////////////////////////////

var (
	alpha26LowerCaseSorted = []string{
		"a", "b", "c", "d", "e",
		"f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o",
		"p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y",
		"z"}

	alpha01MixedCaseSorted = []string{
		"A", "b", "c", "d", "e",
		"f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o",
		"p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y",
		"z"}

	alpha05MixedCaseSorted = []string{
		"A", "b", "c", "d", "e",
		"F", "g", "h", "i", "j",
		"K", "l", "m", "n", "o",
		"P", "q", "r", "s", "t",
		"U", "v", "w", "x", "y",
		"z"}

	alpha10MixedCaseSorted = []string{
		"A", "b", "c", "d", "E",
		"F", "g", "h", "i", "J",
		"K", "l", "m", "n", "O",
		"P", "q", "r", "s", "T",
		"U", "v", "w", "x", "Y",
		"z"}
)

const loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Nullam vestibulum porttitor placerat. Praesent ac lorem mauris. Pellentesque
ultrices, nibh dapibus luctus vulputate, erat ligula lacinia massa, nec
pulvinar elit urna in massa. Interdum et malesuada fames ac ante ipsum primis
in faucibus. Vestibulum ut felis vel turpis interdum interdum. Aenean semper
tempus mattis. Nulla justo libero, pharetra et nulla in, euismod faucibus
dolor. Quisque fringilla, nibh in sollicitudin facilisis, lectus odio semper
nisi, non aliquet nunc arcu non nulla. Etiam facilisis libero vel libero
viverra, vel vulputate neque congue. Ut fermentum quam eget nunc sollicitudin
auctor.

Nam rhoncus imperdiet interdum. Vestibulum facilisis odio dictum velit
condimentum, eu congue nibh gravida. Nulla rutrum eros porttitor eros
suscipit, a vulputate nulla semper. Praesent molestie sollicitudin tincidunt.
Phasellus venenatis mattis mauris, sed fermentum libero ornare quis. Curabitur
mauris odio, posuere ac tincidunt at, posuere a tortor. Pellentesque volutpat
erat ac maximus maximus. Mauris dolor sapien, aliquet non hendrerit eu,
ultrices non arcu. Quisque faucibus eros viverra lacus consectetur iaculis
quis ac ex. In vitae pretium risus, quis fermentum lectus. Etiam vestibulum,
odio in laoreet dapibus, ex nisi pellentesque arcu, vel mattis tortor turpis
nec arcu. Cras scelerisque quam vitae convallis vestibulum. Donec suscipit
urna odio, at luctus libero cursus interdum. Cras mollis bibendum auctor.

Cras tempor, lacus nec sodales facilisis, nisl enim consequat massa, at
efficitur lectus metus sed augue. Nulla vitae risus lacinia, mattis ex vel,
vulputate dui. Nunc id sollicitudin nisl. Aenean a sapien elit. Morbi eget
neque vulputate, consectetur augue eget, consectetur turpis. Curabitur porta,
purus a ultrices bibendum, mauris ipsum posuere dui, at gravida tortor augue
in velit. In mollis enim eu malesuada tristique. Ut dui lectus, aliquet et
felis vel, tristique ultrices nunc. Nunc luctus iaculis ipsum at luctus. Donec
maximus tellus id justo cursus, eu rhoncus diam aliquet. Nam sit amet hendrerit
massa. Fusce enim libero, tincidunt eget nisi at, viverra lobortis enim.

Ut eleifend, nunc in molestie mollis, erat urna interdum neque, id sagittis
ex nisi at libero. Vivamus auctor leo augue, eget ultricies elit dictum eget.
Nam lacinia, urna vel vestibulum placerat, nisl eros accumsan eros, et dictum
nisl nibh ut nisl. Maecenas nec lobortis libero. Vestibulum non vestibulum
quam. Nulla feugiat dignissim augue, eget faucibus metus porta a. Pellentesque
habitant morbi tristique senectus et netus et malesuada fames ac turpis
egestas. Pellentesque lobortis suscipit aliquet. Aenean mattis est velit,
sed consequat lectus ultrices id. Phasellus vel massa ut quam accumsan
pellentesque. Nullam id porttitor orci. Suspendisse suscipit hendrerit tempor.

Mauris finibus lorem nibh, et tristique nunc consequat et. Maecenas sodales
dolor vitae dui fringilla auctor. Pellentesque vitae tempus diam, sed bibendum
risus. Sed volutpat, ligula id iaculis hendrerit, ex eros dapibus risus, ut
scelerisque tellus ipsum ac tellus. Nulla facilisi. Suspendisse ut
pellentesque ex. Phasellus molestie est ac accumsan egestas. Nullam dignissim,
sapien id tempor tincidunt, leo odio dictum ligula, vel lacinia metus turpis
sit amet est. Curabitur mattis dignissim ipsum nec ullamcorper. Proin sagittis
sem arcu, a auctor neque ultrices tincidunt. Etiam luctus elementum pulvinar.
Donec ornare, mauris eget porttitor maximus, dolor turpis rhoncus felis, quis
pretium mi nibh id justo. Proin id pharetra est. Donec auctor tortor eu metus
vehicula, ac hendrerit dui porta. Sed tellus purus, vestibulum a neque in,
eleifend tempor arcu.`

////////////////////////////////////////////////////////////////////////////////
//                    2 CHARS - SINGLE CASE - SORTED                          //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort______2_Chars____LowerCase_Sorted(b *testing.B) {
	benchFoldedSort(b, []string{"a", "b"})
}
func Benchmark_LCasedSort______2_Chars____LowerCase_Sorted(b *testing.B) {
	benchLCasedSort(b, []string{"a", "b"})
}

////////////////////////////////////////////////////////////////////////////////
//                    02 CHARS - SINGLE CASE - SHUFFLED                       //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort______2_Chars____LowerCase_Shuffled(b *testing.B) {
	benchFoldedSort(b, []string{"b", "a"})
}
func Benchmark_LCasedSort______2_Chars____LowerCase_Shuffled(b *testing.B) {
	benchLCasedSort(b, []string{"b", "a"})
}

////////////////////////////////////////////////////////////////////////////////
//                    02 CHARS - 01 MIXED CASE - SORTED                       //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort______2_Chars____MixedCase_Sorted(b *testing.B) {
	benchFoldedSort(b, []string{"A", "b"})
}
func Benchmark_LCasedSort______2_Chars____MixedCase_Sorted(b *testing.B) {
	benchLCasedSort(b, []string{"A", "b"})
}

////////////////////////////////////////////////////////////////////////////////
//                    02 CHARS - 01 MIXED CASE - SHUFFLED                      //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort______2_Chars__1_MixedCase_Shuffled(b *testing.B) {
	benchFoldedSort(b, []string{"b", "A"})
}
func Benchmark_LCasedSort______2_Chars__1_MixedCase_Shuffled(b *testing.B) {
	benchLCasedSort(b, []string{"b", "A"})
}

////////////////////////////////////////////////////////////////////////////////
//                    26 CHARS - SINGLE CASE - SORTED                         //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort_____26_Chars____LowerCase_Sorted(b *testing.B) {
	benchFoldedSort(b, alpha26LowerCaseSorted)
}
func Benchmark_LCasedSort_____26_Chars____LowerCase_Sorted(b *testing.B) {
	benchLCasedSort(b, alpha26LowerCaseSorted)
}

////////////////////////////////////////////////////////////////////////////////
//                   26 CHARS - SINGLE CASE - SHUFFLED                        //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort_____26_Chars____LowerCase_Shuffled(b *testing.B) {
	benchFoldedSort(b, shuffle(copyFrom(alpha26LowerCaseSorted)))
}
func Benchmark_LCasedSort_____26_Chars____LowerCase_Shuffled(b *testing.B) {
	benchLCasedSort(b, shuffle(copyFrom(alpha26LowerCaseSorted)))
}

////////////////////////////////////////////////////////////////////////////////
//                    26 CHARS - 01 MIXED CASE - SORTED                       //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort_____26_Chars__1_MixedCase_Sorted(b *testing.B) {
	benchFoldedSort(b, alpha01MixedCaseSorted)
}
func Benchmark_LCasedSort_____26_Chars__1_MixedCase_Sorted(b *testing.B) {
	benchLCasedSort(b, alpha01MixedCaseSorted)
}

////////////////////////////////////////////////////////////////////////////////
//                    26 CHARS - 01 MIXED CASE - SHUFFLED                     //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort_____26_Chars__1_MixedCase_Shuffled(b *testing.B) {
	benchFoldedSort(b, shuffle(copyFrom(alpha01MixedCaseSorted)))
}
func Benchmark_LCasedSort_____26_Chars__1_MixedCase_Shuffled(b *testing.B) {
	benchLCasedSort(b, shuffle(copyFrom(alpha01MixedCaseSorted)))
}

////////////////////////////////////////////////////////////////////////////////
//                    26 CHARS - 05 MIXED CASE - SORTED                       //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort_____26_Chars__5_MixedCase_Sorted(b *testing.B) {
	benchFoldedSort(b, alpha05MixedCaseSorted)
}
func Benchmark_LCasedSort_____26_Chars__5_MixedCase_Sorted(b *testing.B) {
	benchLCasedSort(b, alpha05MixedCaseSorted)
}

////////////////////////////////////////////////////////////////////////////////
//                   26 CHARS - 05 MIXED CASE - SHUFFLED                      //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort_____26_Chars__5_MixedCase_Shuffled(b *testing.B) {
	benchFoldedSort(b, shuffle(copyFrom(alpha05MixedCaseSorted)))
}
func Benchmark_LCasedSort_____26_Chars__5_MixedCase_Shuffled(b *testing.B) {
	benchLCasedSort(b, shuffle(copyFrom(alpha05MixedCaseSorted)))
}

////////////////////////////////////////////////////////////////////////////////
//                    26 CHARS - 10 MIXED CASE - SORTED                       //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort_____26_Chars_10_MixedCase_Sorted(b *testing.B) {
	benchFoldedSort(b, alpha10MixedCaseSorted)
}
func Benchmark_LCasedSort_____26_Chars_10_MixedCase_Sorted(b *testing.B) {
	benchLCasedSort(b, alpha10MixedCaseSorted)
}

////////////////////////////////////////////////////////////////////////////////
//                   26 CHARS - 10 MIXED CASE - SHUFFLED                      //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort_____26_Chars_10_MixedCase_Shuffled(b *testing.B) {
	benchFoldedSort(b, shuffle(copyFrom(alpha10MixedCaseSorted)))
}
func Benchmark_LCasedSort_____26_Chars_10_MixedCase_Shuffled(b *testing.B) {
	benchLCasedSort(b, shuffle(copyFrom(alpha10MixedCaseSorted)))
}

////////////////////////////////////////////////////////////////////////////////
//                 LOREM IPSUM - 542 WORDS - MIXED CASE - SORTED              //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort____542_Words____MixedCase_Sorted(b *testing.B) {
	data := strings.Fields(loremIpsum)
	sort.Strings(data)
	benchFoldedSort(b, data)
}
func Benchmark_LCasedSort____542_Words____MixedCase_Sorted(b *testing.B) {
	data := strings.Fields(loremIpsum)
	sort.Strings(data)
	benchLCasedSort(b, data)
}

////////////////////////////////////////////////////////////////////////////////
//                 LOREM IPSUM - 542 WORDS - MIXED CASE - SHUFFLED            //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort____542_Words____MixedCase_Shuffled(b *testing.B) {
	benchFoldedSort(b, strings.Fields(loremIpsum))
}
func Benchmark_LCasedSort____542_Words____MixedCase_Shuffled(b *testing.B) {
	benchLCasedSort(b, strings.Fields(loremIpsum))
}

////////////////////////////////////////////////////////////////////////////////
//                 LOREM IPSUM - 54,200 WORDS - MIXED CASE                    //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort__54200_Words____MixedCase_Shuffled(b *testing.B) {
	benchFoldedSort(b, duplicate(100, strings.Fields(loremIpsum)))
}

func Benchmark_LCasedSort__54200_Words____MixedCase_Shuffled(b *testing.B) {
	benchLCasedSort(b, duplicate(100, strings.Fields(loremIpsum)))
}

////////////////////////////////////////////////////////////////////////////////
//                LOREM IPSUM - 542,000 WORDS - MIXED CASE                    //
////////////////////////////////////////////////////////////////////////////////
func Benchmark_FoldedSort_542000_Words____MixedCase_Shuffled(b *testing.B) {
	benchFoldedSort(b, duplicate(1000, strings.Fields(loremIpsum)))
}

func Benchmark_LCasedSort_542000_Words____MixedCase_Shuffled(b *testing.B) {
	benchLCasedSort(b, duplicate(1000, strings.Fields(loremIpsum)))
}

////////////////////////////////////////////////////////////////////////////////
//                     LCASE-INSENSITIVE SORT IFACE IMPL                      //
////////////////////////////////////////////////////////////////////////////////

type toLower []string

func (s toLower) Len() int {
	return len(s)
}

func (s toLower) Less(i, j int) bool {
	return strings.ToLower(s[i]) < strings.ToLower(s[j])
}

func (s toLower) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

////////////////////////////////////////////////////////////////////////////////
//                                   UTILS                                    //
////////////////////////////////////////////////////////////////////////////////
func duplicate(count int, in []string) []string {
	out := make([]string, len(in)*count)
	for i := 0; i < len(in); i++ {
		for j := 0; j < count; j++ {
			out[i+(j*len(in))] = in[i]
		}
	}
	return out
}

func copyFrom(src []string) []string {
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

func shuffle(src []string) []string {
	for i := range src {
		j := rand.Intn(i + 1)
		src[i], src[j] = src[j], src[i]
	}
	return src
}

func benchFoldedSort(b *testing.B, data []string) {
	benchSort(b, newFoldedSortWrapper, data)
}

func benchLCasedSort(b *testing.B, data []string) {
	benchSort(b, newLCasedSortWrapper, data)
}

func newFoldedSortWrapper(src []string) sort.Interface {
	return sort.Interface(sortfold.StringSlice(copyFrom(src)))
}

func newLCasedSortWrapper(src []string) sort.Interface {
	return sort.Interface(toLower(copyFrom(src)))
}

func benchSort(b *testing.B, f func([]string) sort.Interface, d []string) {
	data := make([]sort.Interface, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = f(d)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Sort(data[i])
	}
}
