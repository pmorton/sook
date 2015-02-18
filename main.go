package main

import (
  "flag"
  "text/template"
  "os"
  "fmt"
  "strings"
  "io"
  "log"
  )

var inFile string;
var outFile string;
var t *template.Template;

func init() {
  flag.StringVar(&inFile, "in", "", "Input template file")
  flag.StringVar(&outFile, "out", "", "Output template file")
  flag.Parse()
  t = template.Must(template.ParseFiles(inFile))
}

func main() {
  var writer io.WriteCloser

  if outFile == "" {
    writer = os.Stdout
  } else {
    var err error;
    writer, err = os.Create(outFile)
    defer writer.Close()
    if err != nil {
      log.Fatalf(fmt.Sprintf("%s",err))
    }

  }

  err := t.ExecuteTemplate(writer, inFile, getEnvironment())
  if err != nil {
    fmt.Printf("%s",err)
    os.Exit(1)
  }
}


func getEnvironment() (map[string]string) {
  env := os.Environ()
  rtn := make(map[string]string)
  for _,item := range env {
    splits := strings.Split(item, "=")
    rtn[splits[0]]=splits[1]
  }
  return rtn
}
