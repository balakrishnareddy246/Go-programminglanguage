package main 

import "fmt"

func main () {

  var name string ="Balakrishna reddy"

  const pi float64 = 3.1412435
  win := true 
  x :=5
  

// fmt.Println(len(name))
// fmt.Println(name + " os a chill dude")


  fmt.Printf("%.3f \n",pi)
  fmt.Printf("%T \n", name)
  fmt.Printf("%t \n", win)
  fmt.Printf("%d \n", x)
  fmt.Printf("%b \n", 25)
  fmt.Printf("%c \n", 34)
  fmt.Printf("%x \n", 15)
}