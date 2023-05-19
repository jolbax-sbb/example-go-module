package example_go_module

import "fmt"
import "rsc.io/quote"

// Hello is a function that prints a some useful text
func Hello() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())
}
