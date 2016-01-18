package main

import "fmt"

func main() {
  //channel默认上是阻塞的，也就是说，如果Channel满了，就阻塞写，如果Channel空了，就阻塞读。
  //阻塞的含义就是一直等到轮到它为止。单有时候我们会收到 fatal error: all goroutines are asleep - deadlock!
  messages := make(chan string, 1)
  messages <- "ping"
  // messages <- "hello"
  // _Send_ a value into a channel using the `channel <-`
  // syntax. Here we send `"ping"`  to the `messages`
  // channel we made above, from a new goroutine.
  go func() { messages <- "hi" }()

  // The `<-channel` syntax _receives_ a value from the
  // channel. Here we'll receive the `"ping"` message
  // we sent above and print it out.

  // go func() {messages <- "visit google"} ()
  msg := <-messages
  msg = <-messages
  fmt.Println(msg)
}
