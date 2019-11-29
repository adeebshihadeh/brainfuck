package main

import(
  "os"
  "fmt"
  "io/ioutil"
)

func main() {
  var mem [30000]int8
  var ptr uint8

  prog, _ := ioutil.ReadFile(os.Args[1])

  pc := 0
  for pc < len(prog) {
    c := prog[pc]
    pc++

    switch c {
    case '>':
      ptr += 1
    case '<':
      ptr -= 1
    case '+':
      mem[ptr] += 1
    case '-':
      mem[ptr] -= 1
    case '[':
      if mem[ptr] == 0 {
        cnt := 0
        for pc < len(prog) {
          if prog[pc] == ']' {
            if cnt == 0 {
              pc++
              break
            }
            cnt--
          } else if prog[pc] == '[' {
            cnt++
          }
          pc++
        }
      }
    case ']':
      if mem[ptr] != 0 {
        cnt := 0
        pc -= 2
        for pc > 0 {
          if prog[pc] == '[' {
            if cnt == 0 {
              break
            }
            cnt -= 1
          } else if prog[pc] == ']' {
            cnt++
          }
          pc--
        }
      }
    case '.':
      fmt.Print(string(mem[ptr]))
    case ',':
      fmt.Println("TODO: implement input")
    }
  }
  fmt.Println()
}

