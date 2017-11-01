package main

import (
	"fmt"
	"math"
)


// if you are using Pointer in Receiver function then you can mutate the original state of struct

type rect struct {
	width, height float64
}

type cir struct {
	radius float64
}

func (r *rect ) area() float64 {
	//r.height=1222 if you change the height then original height will get changed  and in future calls heights will be 122
  return r.width * r.height
}

func (c *cir) area() float64{
 return math.Pi*c.radius
}

func main() {
	r := rect{3,4}
	fmt.Println("Rectangle" , r.area() )
	c:= cir{3}
	fmt.Println("Circle" , 	c.area())
}
