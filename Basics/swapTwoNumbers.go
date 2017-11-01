package main

import "fmt"

 //Using Pointer
func main() {
	a:=10
	b:=20
	swap(&a,&b)
	fmt.Println(a)
	fmt.Println(b)
}

func swap(a *int ,b *int){
		t := new(int)

		 *t = *a
		 *a=*b
	 	 *b= *t


}

//Using Shortcut
/*
func main()  {
	a:= 10
	b:= 20
	fmt.Println(a)
	fmt.Println(b)
	a,b=b,a


	fmt.Println(a)
	fmt.Println(b)
}
*/


