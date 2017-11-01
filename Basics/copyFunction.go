package main

import "fmt"

//Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).
/// when using Copy function it changes the Reference of the variable and points it to source reference
func main() {
	a := [] int{1,2,3,4}
	b := [] int {6,7,8,9}
	//Here, the check variable is used to hold a reference to the original slice description to make sure it is really copied.
	check := a
	fmt.Println(check)
	copy(a,b)
	fmt.Print(a)
	fmt.Print(b)
  	fmt.Println(check)
}
// when using assignment operator it makes the copy of value of assigned variable .
//func main()  {
//	a := []int{1, 2}
//	b := []int{3, 4}
//	check := a
//	a = b
//	fmt.Println(a, b, check)
//
//}