/*
const (
    C1 = iota + 1
    //todo: skip one
    C3
    C4
)
fmt.Println(C1, C3, C4) // "1 3 4"
*/
/*
func SumProduct(a, b int) (int, int) {
    return a+b, a*b
}

func main() {
    // I only want the sum, but not the product
   //todo: skip the second return value
   //todo: do i need to do this?
   //todo: where do I use this most?
     // the product gets discarded
    fmt.Println(sum) // prints 3
}
*/
/*
   pets := []string{"dog", "cat", "fish"}

   // Range returns both the current index and value
   // but sometimes you may only want to use the value
   for _, pet := range pets {
       fmt.Println(pet)
   }
*/

//in half-written program
//Silence the compiler
//It can be used to during development to avoid compiler errors about unused imports and variables in a half-written program.

package main

import (
	"log"
	"os"
)

//todo: fix my partially written code

func main() {
	f, err := os.Open("gocon.go")
	if err != nil {
		log.Fatal(err)
	}
	// todo: reference file
}
