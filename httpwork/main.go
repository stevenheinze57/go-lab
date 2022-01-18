
package main

import "fmt"
import "os"
import "net/http"
import "io"

type logWriter struct{}

func main() {
  resp, err := http.Get("https://google.com")

  if (err != nil) {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }

  // Body which implement the Read interface can be called passing in a byte slice object
  // That data in the body is then saved to the byte slie object which can then be printed
  // This is essentially importing data into our program from some source
  bs := make([]byte, 99999)
  resp.Body.Read(bs)
  // fmt.Println(string(bs))

  // Taking a source of data, converting it, and writing it to an output
  // The Copy function is taking in something that implements the Writer interface as the first arg
  // and something that implements the Reader interface as the second arg
  // io.Copy(os.Stdout, resp.Body)

  // Calling the custom implementation of our Write interface method
  lw := logWriter{}
  io.Copy(lw, resp.Body)
}

// To implement the Write interface, the method name of this receiver function must match the capitalized function "Write"
func (logWriter) Write(s []byte) (n int, e error) {
  fmt.Println(string(s))
  fmt.Println("Just wrote this many bytes: ", string(len(s)))
  return len(s), nil
}
