package main

import (
  "fmt"
  "time"
)

func main() {
  s1 := constructor()
  s1.insert("one", "1")
  s1.insert("two", "2")
  s1.insert("three", "4")
  fmt.Printf("%v is %v!\n", "three", s1.m["three"])
  println("wrong, gotta fix...")
  time.Sleep(3 * time.Second)
  s1.update("three", "3")
  fmt.Printf("%v is %v!\n", "three", s1.m["three"])
  println(s1.contains("three")) // true
  s1.delete("three")
  println(s1.contains("three")) // false
}


type store struct {
  m map[string]string
}

func constructor() store {
  return store{
    m: map[string]string{},
  }
}

func (s store) insert(k, v string) {
  if s.m[k] == "" {
    s.m[k] = v
  } else {
    println(k, "already exists")
  }
}

func (s store) update(k, v string) {
  if s.m[k] != "" {
    s.m[k] = v
  } else {
    println(k, "doesn't exist")
  }
  // if _, ok := s.m[k]; ok {
		// s.m[k] = v
	// }
}

func (s store) upsert(k, v string) {
  s.m[k] = v
}

func (s store) contains(k string) bool {
  if s.m[k] != "" {
    return true
  } else {
    return false
  }
}

func (s store) delete(k string) {
  delete(s.m, k)
}
