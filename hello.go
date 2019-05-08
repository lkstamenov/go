package main

import (
   "fmt"

)

func main () {
  // STDOUT
   fmt.Println("Hello World!")
  // Variables
   x := 1
   y := 12
   sum := x + y
   fmt.Println("sum is: ", sum)
  // if else statement
   if sum > 20 {

     fmt.Println("sum is greater than 20")
   } else if x < 2 {
      fmt.Println("x is less than 2")
   } else  {
   }
  // Array
   var a [5]int
   fmt.Println(a)
  // Slice
   slice := []int{5,7,90,12,8}
   slice = append(slice, 44)
   fmt.Println(slice)
  // Map (Dictionary)
   vertices := make(map[string]int)
   vertices["triangle"] = 2
   vertices["square"] = 3
   vertices["dodecagon"] = 12
   fmt.Println(vertices)
   fmt.Println(vertices["triangle"])
   delete(vertices, "square")
   fmt.Println(vertices)

}
