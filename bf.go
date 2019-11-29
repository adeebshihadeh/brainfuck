package main

import(
  "os"
  "fmt"
  "strings"
  "io/ioutil"
)

func main() {
  // TODO: max amount of memory cells?
  var memory [1000]uint32
  var ptr uint8

  prog, _ := ioutil.ReadFile(os.Args[1])

  pos := 0
  for pos < len(prog) {
    b := prog[pos]
    pos += 1

    switch b {
    case '>':
      ptr += 1
    case '<':
      ptr -= 1
    case '+':
      memory[ptr] += 1
    case '-':
      memory[ptr] -= 1
    case '[':
      if memory[ptr] == 0 {
        pos = strings.IndexByte(string(prog[pos:]), ']') + 1
      }
    case ']':
      if memory[ptr] != 0 {
        pos = strings.LastIndexByte(string(prog[:pos]), '[') + 1
      }
    case '.':
      fmt.Print(string(memory[ptr]))
    }
  }
  fmt.Println()
}

