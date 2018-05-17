package main

import "fmt"

func main() {
  m := map[int]int{}
  for i := 1; i <= 100; i++ {
    m[i] = i * i
  }
  sumEven, sumOdd := 0, 0
  for k, v := range m {
    if k % 2 == 0 {
      sumEven += v
    } else {
      sumOdd += v
    }
  }
  fmt.Printf("The sum of squares for even numbers is %v \n", sumEven)
  fmt.Printf("The sum of squares for odd numbers is %v \n", sumOdd)
}
