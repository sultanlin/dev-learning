package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
  for i:=0; i<5; i++ {
    fmt.Println(text)
    time.Sleep(time.Second * 1)
  }
  channel <- "Done"
}

func main() {
  // var channel chan string
  channel := make(chan string)

  go printMessage("Go is great", channel)
  <- channel
}
