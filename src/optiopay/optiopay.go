package main

import (
	"fmt"
	"optiopay/lib/lca"
)

/*type data struct {
	name string
	children []*data
}
*/

func main() {
	var one, two, three, four, five, six, seven *lca.Data
	one = &lca.Data{Name: "1"}
	two = &lca.Data{Name: "2"}
	three = &lca.Data{Name: "3"}
	four = &lca.Data{Name: "4"}
	five = &lca.Data{Name: "5"}
	six = &lca.Data{Name: "6"}
	seven = &lca.Data{Name: "7"}

	one.Children = append(one.Children, two)
	one.Children = append(one.Children, three)
	two.Children = append(two.Children, four)
	two.Children = append(two.Children, five)
	three.Children = append(three.Children, six)
	three.Children = append(three.Children, seven)

	lca.Init(one)
	fmt.Printf("%s", lca.Search("4","5"))
}
