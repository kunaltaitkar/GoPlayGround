package subPck

import (
	"strings"
	"testing"
)

// func main() {

// }

func concatWithStringBuilder() {
	a := []string{"asd", "dsdsd", "sdsdsd"}
	var str strings.Builder

	for _, s := range a {
		str.WriteString(s)
	}
	// fmt.Println("result:", str)
}

func concatWithoutStringBuilder() {
	a := []string{"asd", "dsdsd", "sdsdsd"}
	str := ""

	for _, s := range a {
		str += s
	}
	// strconv.Ito
	// reflect.TypeOf()

	// fmt.Println("result2:", str)
}

func BenchmarkConcatWithStringBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatWithStringBuilder()
	}
}

func BenchmarkConcatWithoutStringBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatWithoutStringBuilder()
	}
}
