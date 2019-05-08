package main

import (
   "fmt"
   "errors"
   "math"
)


// Functions

func suma(x int, y int) int {
  return x + y
}

func sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, errors.New("Undefined for negative number")
  }
  return math.Sqrt(x), nil
}

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

  // Loops
  // for loop

  for i :=0; i < 5; i++ { 
    fmt.Println(i)
  }
  // while loop
  i := 0
  for i < 8 {
    fmt.Println(i)
    i++
  }
  // Loop through an array
  arr := []string{"a", "b", "c"}

  for index, value := range arr {
     fmt.Println("index", index, "value", value)
  }
  // Loop and maps
  m := make(map[string]string)
  m["a"] = "alpha"
  m["b"] = "beta"
  for k, value := range m {
    fmt.Println("key: ", k, "value:", value)
  
  } 
  // Functions
  result := suma(3,5)
  fmt.Println(result)
  squareroot, errorvalue := sqrt(4)
  fmt.Println(squareroot, errorvalue)
}


