
package main

import "testing"
import "fmt"
import "os"

func TestNewDeck(t *testing.T)  {
  d := newDeck()

  if (len(d) != 16) {
    t.Errorf("Expected length of 16 on newDeck but received %v", len(d))
  } else {
    // Pass -v to go test command to have this output printed
    fmt.Println("Deck length test passed")
  }

  if (d[0] != "Ace of Spades") {
    t.Errorf("Expected 'Ace of Spades' but got %v", d[0])
  }
}

func TestFirstCardInNewDeck(t *testing.T) {
  d := newDeck()

  if (d[0] != "Ace of Spades") {
    t.Errorf("Expected 'Ace of Spades' but got %v", d[0])
  } else {
    fmt.Println("First card in deck test passed")
  }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T)  {
  decktestfile := "_decktesting.txt"
  os.Remove(decktestfile)

  deck := newDeck()
  deck.saveToFile(decktestfile)

  loadedDeck := newDeckFromFile(decktestfile)

  if (len(loadedDeck) != 16) {
    t.Errorf("Expected 16 cards in loaded deck but got %v", len(loadedDeck))
  }

  os.Remove(decktestfile)
}
