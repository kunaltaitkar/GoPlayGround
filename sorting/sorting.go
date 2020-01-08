package main

import (
	"fmt"
	"sort"
)

func main() {

	dict1 := []string{"Coders", "loves", "the", "algorithms"}

	//To check lexically sorted or not
	fmt.Println("Lexically Sorted? ", sort.StringsAreSorted(dict1))

	//To sort in lexical order
	dict2 := sort.StringSlice{"Coders", "loves", "the", "algorithms"}
	dict2.Sort()
	fmt.Println("After Sorting ", dict2)

}
