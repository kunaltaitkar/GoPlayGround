package main

import "fmt"

type Dense interface {
	Density() int
}

type Rock struct {
	Volume int
	Mass   int
}

type Geode struct{}

func (r Rock) Density() int {
	return r.Mass / r.Volume
}

func (g Geode) Density() int {
	return 100
}

func IsItDenser(r, g Dense) bool {
	return r.Density() > g.Density()
}

func main() {
	r := Rock{10, 20}
	g := Geode{}
	var a interface{}
	a = 10
	fmt.Println("OUTPUT:", IsItDenser(r, g))

	switch str := a.(type) {
	case string:
		fmt.Println("this is string", str)
	default:
		fmt.Println("this is default")
	}

}
