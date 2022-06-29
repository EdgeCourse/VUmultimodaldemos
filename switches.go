//https://go.dev/play/p/PDxrUoha6rf
//Switch statements express conditionals across many branches.
package main

func main() {
	/*
		//Here’s a basic switch.
		i := 2
		fmt.Print("Write ", i, " as ")
		switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		}
	*/
	//You can use commas to separate multiple expressions in the same case statement. We use the optional default case in this example as well.
	/*
	   switch time.Now().Weekday() {
	       case time.Saturday, time.Sunday:
	           fmt.Println("It's the weekend")
	       default:
	           fmt.Println("It's a weekday")
	       }
	*/
	//switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.
	//todo: what is this called?
	/*
	   t := time.Now()
	       switch {
	       case t.Hour() < 12:
	           fmt.Println("It's before noon")
	       default:
	           fmt.Println("It's after noon")
	       }
	*/

	//todo: what is this called?
	//A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause.

	/*
	   whatAmI := func(i interface{}) {
	           switch t := i.(type) {
	           case bool:
	               fmt.Println("I'm a bool")
	           case int:
	               fmt.Println("I'm an int")
	           default:
	               fmt.Printf("Don't know type %T\n", t)
	           }
	       }
	       whatAmI(true)
	       whatAmI(1)
	       whatAmI("hey")
	   }
	*/
	/*
		//todo: what is the behavior, how to fallthrough?
			   i := 45
			   switch {
			   case i < 10:
			   	fmt.Println("i is less than 10")

			   case i < 50:
			   	fmt.Println("i is less than 50")

			   case i < 100:
			   	fmt.Println("i is less than 100")
			   }
	*/
	/*
	   Output

	   i is less than 50
	   i is less than 100
	   fallthrough needs to be final statement within the switch block. If it is not then compiler raise error

	   fallthrough statement out of place
	*/

	//todo: can we use break?
	/*
	   switch char := "b"; char {
	   case "a":
	       fmt.Println("a")
	   case "b":
	       fmt.Println("b")

	       fmt.Println("after b")
	   default:
	       fmt.Println("No matching character")
	   }
	*/
}
