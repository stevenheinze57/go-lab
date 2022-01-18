
package main

import "fmt"

func main()  {
  evenAndOddNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  for _, num := range evenAndOddNumbers {
    if ((num % 2) > 0) {
      fmt.Println("Number ", num, " is odd.")
    } else {
      fmt.Println("Number ", num, " is even")
    }
  }
}
