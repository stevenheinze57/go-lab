
package main

import "fmt"

type bot interface {
  // Get greeting method takes no arguments and has a string return type
  getGreeting() string
}

// Interfaces are implicit, so there is no need to defined explicit relationships between an interface and an object
type englishBot struct{}
type spanishBot struct{}

func main() {
  eb := englishBot{}
  sb := spanishBot{}

  printGreeting(eb)
  printGreeting(sb)
}

// the bot object must implement the interface
func printGreeting(b bot) {
  fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
  return "Hello..."
}

func (spanishBot) getGreeting() string {
  return "Hola..."
}
