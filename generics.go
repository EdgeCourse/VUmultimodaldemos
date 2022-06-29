package main

import (
	"fmt"
)
//todo: make a pure function
func PrintSlice[) {
	fmt.Println("Generics")
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func main() {
	PrintSlice([]int{1, 2, 3})
	PrintSlice([]string{"a", "b", "c"})
	PrintSlice([]float64{1.2, -2.33, 4.55})

}

//numeric
/*
package main

import (
	"fmt"
)

//todo: name your interface
type  interface {
	//todo: make numeric
	//~int //anything derived from int
}

//todo: use custom generic
func Add[](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("4 + 3 =", Add(4, 3))
	fmt.Println("4.1 + 3.2 =", Add(4.1, 3.2))
}
*/
