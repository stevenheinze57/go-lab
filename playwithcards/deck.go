package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "os"
  "math/rand"
  "time"
)

// Create a new type named deck which is defined as a slice of strings
type deck []string

// Any variable of type deck gets access to this function (this is refered to as a receiver)
// a receiver function will be available to any instance of the receiver type
func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}

func (d deck) toString() string {
  return strings.Join([]string(d), ",")
}

// returns an error if one occurs
func (d deck) saveToFile(filename string) error {
  return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
  // random generator requires a source for randomization
  source := rand.NewSource(time.Now().UnixNano())
  r := rand.New(source)

  for i := range d {
    newPosition := r.Intn(len(d) - 1)
    // multi-value assignment here
    d[i], d[newPosition] = d[newPosition], d[i]
  }
}

func newDeckFromFile(filename string) deck {
  // demonstrated that the ReadFile function is returning two objects (a byte slice, and an error if it occurs or nil)
  bs, err := ioutil.ReadFile(filename)
  if (err != nil) {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }

  // typecasting the byte slice to a string and splitting it by each comma
  s := strings.Split(string(bs), ",")
  return deck(s)
}

func newDeck() deck {
  cards := deck{}

  cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
  cardValues := []string{"Ace", "Two", "Three", "Four"}

  for _, suit := range cardSuits {
    for _, value := range cardValues {
      cards = append(cards, (value + " of " + suit))
    }
  }

  return cards
}

// returning two separate objects
func deal(d deck, handSize int) (deck, deck) {
  return d[:handSize], d[handSize:]
}

