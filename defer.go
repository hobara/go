package main

import "time"

func main() {
  defer println("🍜🍜🍜")
  defer println("Ramen is ready!")
  println("...cooking...")
  time.Sleep(3 * time.Second)
}
