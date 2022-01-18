
package main

import "fmt"
import "net/http"
import "time"

// Concurrency is not parallelism: a go routine with one core will run multiple threads one at a time, whereas with parallelism
// with multiple cores you can have multiple threads executing in parallel. The go sheduler for routines will manage this
// amongst the number fo cores available.

func main() {
  links := []string{
    "https://google.com",
    "https://facebook.com",
    "https://stackoverflow.com",
    "https://golang.org",
    "https://amazon.com",
  }

  // Create a channel of type string
  c := make(chan string)

  fmt.Println(">>> Trying first for loop <<<")
  for _, link := range links {
    // Startup the go routine thread
    go checkLink(link, c)
  }
  // Wait for messages in the channel and logs them, but will only print the first set of data
  fmt.Println(<- c)
  fmt.Println(<- c)
  fmt.Println(<- c)
  fmt.Println(<- c)
  fmt.Println(<- c)

  fmt.Println(">>> Trying second for loop <<<")
  for _, link := range links {
    go checkLink(link, c)
  }
  for i := 0; i < len(links); i++ {
    fmt.Println(<- c)
  }

  fmt.Println(">>> Trying third for loop <<<")
  for _, link := range links {
    time.Sleep(5 * time.Second)
    go checkLink(link, c)
  }
  // for infinity loop
  // for {
  //   // first argument is any and all values that are being received in the channel
  //   // this causes a recursive loop
  //   go checkLink(<- c, c)
  // }

  // An alternative syntax to the infinite for loop above
  // wait until the channel has received some value, then assign it to the value "l" and execute the loop
  // for l := range c {
  //   go checkLink(l, c)
  // }

  // Call to a function literal to execute
  // Since golang is a pass-by-value language we want to pass in the value of the string variable
  // to the function literal. This ensures the value is used and not the memory address reference 
  for l := range c {
    go func(link string) {
      time.Sleep(5 * time.Second)
      checkLink(link, c)
    }(l)
  }
}

func checkLink(link string, c chan string) {
  _, err := http.Get(link)
  if (err != nil) {
    fmt.Println(link, "might be down!")
    // Send message into channel
    c <- link
    return
  }

  fmt.Println(link, "is up!")
  // Send message into channel
  c <- link
}
