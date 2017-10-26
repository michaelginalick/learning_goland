package main


import (
  "flag"
  "os"
  "fmt"
  "encoding/csv"
  "strings"
)

func main() {

  //set flag
  csvFileName := flag.String("csv", "problems.csv", "a file in the form of 'question,answer'")
  // parse flag
  flag.Parse()
  //open file
  file, err := os.Open(*csvFileName)

  if err != nil {
    exit(fmt.Sprintf("failed to open file name: %s/n", *csvFileName))
    os.Exit(1)
  }
  //read in file from reader
  reader := csv.NewReader(file)
  // read all lines from reader
  lines, err := reader.ReadAll()

  if err != nil {
    exit("Failed to parse the provided csv file")
  }

  problems := parseLines(lines)

  correct := 0
  for i, p := range problems {
    fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

    var answer string
    fmt.Scanf("%s\n", &answer)

    if answer == p.a {
      correct++
    }
  }

  fmt.Printf("you scored %d out of %d", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
  //make problems from problem struct
  ret := make([]problem, len(lines))

  for i, line := range lines {
    ret[i] = problem{
      q: line[0],
      a: strings.TrimeSpace(line[1]),
    }
  }

  return ret
}

type problem struct {
  q string
  a string
}


func exit(msg string) {
  fmt.Println(msg)
  os.Exit(1)
}