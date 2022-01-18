package main

import (
  "fmt"
)

func main() {
  fmt.Println("Playing with cards....")

  fmt.Println("Obtaining hand....")
  cards := newDeck()
  hand, remainingDeck := deal(cards, 5)

  fmt.Println("Showing hand....")
  hand.print()

  fmt.Println("Showing remaining deck....")
  remainingDeck.print()

  fmt.Println("Printing hand as string....")
  fmt.Println(hand.toString())

  fmt.Println("Writing hand to local file....")
  hand.saveToFile("hand.txt")

  fmt.Println("Writing remaining deck to local file....")
  remainingDeck.saveToFile("remainingDeck.txt")

  fmt.Println("Creating new deck from remaining deck file....")
  newDeck := newDeckFromFile("remainingDeck.txt")

  fmt.Println("Printing new deck....")
  newDeck.print()

  fmt.Println("Shuffling new deck....")
  newDeck.shuffle()

  fmt.Println("Printing new deck after shuffling....")
  newDeck.print()
}
