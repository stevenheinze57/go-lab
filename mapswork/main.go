
package main

import "fmt"

func main() {
  // creates a map where all keys are of string type, and all values are of string type
  colors := map[string]string{
    "red": "0001",
    "green": "0010",
    "white": "0100",
  }
  printMap(colors)

  // make is a builtin function in golang for creating an object of the passed in type
  cars := make(map[int]string)
  cars[0] = "Chevy Lumina"
  cars[1] = "Dodge Avenger"
  cars[2] = "Chevy Silverado"
  fmt.Println(cars)

  // builtin golang function to delete an element from a map
  delete(cars, 0)
  fmt.Println(cars)
}

func printColorsMap (c map[string]string) {
  for color, number := range c {
    fmt.Println("Color: " + color + ", Number: " + number)
  }
}