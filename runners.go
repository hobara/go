package main

import "time"

type runner interface {
  name() string
  run(distance int) (seconds int)
}

func race(distance int, runners ...runner) (winner string) {
  fastest := 0
  for i, runner := range runners {
    res := runner.run(distance)
    println(runner.name(), "ran ---", res, "seconds!")
    time.Sleep(1 * time.Second)
    if i == 0 {
      fastest = res
    } else if res < fastest {
      fastest = res
      winner = "Winner is " + runner.name()
    } else if res == fastest {
      winner = "Tie race! Do it again!"
    }
  }
  return winner
}

func main() {
  var (
    r1 = dog{}
    r2 = cat{}
    r3 = bear{}
    r4 = fox{}
    r5 = human{}
  )
  println(race(100, r1, r2, r3, r4, r5))
}

type dog struct{}

func (dog) name() string {
  return "Lily"
}

func (dog) run(distance int) int {
  return distance / 3
}

type cat struct{}

func (cat) name() string {
  return "Luna"
}

func (cat) run(distance int) int {
  return distance / 4
}

type bear struct{}

func (bear) name() string {
  return "Winne the pooh"
}

func (bear) run(distance int) int {
  return distance / 2
}

type fox struct{}

func (fox) name() string {
  return "Silver"
}

func (fox) run(distance int) int {
  return distance / 3
}

type human struct{}

func (human) name() string {
  return "Hiro"
}

func (human) run(distance int) int {
  return distance / 2
}
