/*
 * Main tests the code witj sample files for a dictionary and a document
 * and provides the user with preferences for the programs operation using
 * command line flags.
 * 14/5/18
 * Author Callan Taylor
 */
package main

import "flag"
import "fmt"
import "io/ioutil"
import "time"
import "log"
import "strings"

var defaultSize int = 10
var unknownWords int
var allWords []string

func main() {

  containerType := flag.String("type", "FLEX_ARRAY", "container type(RBT, FLEX_ARRAY)")
  containerSize := flag.Int("size", defaultSize, "hashTable size")
  dictionaryFile := flag.String("dictionary", "sampleDictionary.txt", "dictionary words file")
  documentFile := flag.String("document", "sampleWords.txt", "text file to spellcheck")
  printHTable := flag.String("print", "No", "print htable only (Yes/No)")
  infoTable := flag.String("info", "no", "show info on table(Yes/No)")
  flag.Parse()

  h := newHtable(*containerSize, *containerType)

  dictionary := fileToSlice(dictionaryFile)
  document := fileToSlice(documentFile)

  startFill := time.Now()
  for i := 0; i < len(dictionary); i++ {
    htableInsertWord(h, dictionary[i])
  }
  fillTime := time.Since(startFill)

  if *printHTable == "Yes" || *printHTable == "yes" {
    htablePrint(h)
  } else {
    start := time.Now()
    for i := 0; i < len(document); i++ {
      if htableSearch(h, document[i]) == false {
        fmt.Println(document[i])
        unknownWords++
      }
    }
    searchTime := time.Since(start)
    if *infoTable == "yes" || *infoTable == "Yes" {
      fmt.Println("\nFill time:", fillTime)
      fmt.Println("Search time:", searchTime)
      fmt.Println("Unknown Words:", unknownWords)
    }
  }
}


func fileToSlice(file *string) []string {

  stream, err := ioutil.ReadFile(*file)

  if err != nil {
    log.Fatal(err)
  }

  s := string(stream)
  s = strings.ToLower(s)
  sentences := strings.Split(s, "\n")
  for i := 0; i < len(sentences); i++ {
    words := strings.Split(sentences[i], " ")
    for j := 0; j < len(words); j++ {
      allWords = append(allWords, words[j])
    }
  }

  for i := 0; i < len(allWords); i++ {
    allWords[i] = strings.Trim(allWords[i], " ")
    allWords[i] = strings.Trim(allWords[i], ".")
  }
  return allWords
}
