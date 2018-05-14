/*
* Implementation of a Hash Table which stores containers of
* either RED_BLACK_TREE or Slice
* 14/5/18
* Author Callan Taylor
*/
package main

import "fmt"


type htable struct {
  keys []*container
  frequencies []int
  numKeys int
  capacity int
}


func newHtable(size int, t string) *htable {
  keyArray := make([]*container, size)
  frequenciesArray := make([]int, size)
  var numkeys int = 0
  var capacity int = size

  for i := 0; i < size; i++ {
    frequenciesArray[i] = 0
    keyArray[i] = newContainer(t)
  }

  newHtable := htable{keys:keyArray, frequencies:frequenciesArray,
    numKeys:numkeys, capacity:capacity}

  return &newHtable
}


func htableWordToInt(word string) uint  {
  var result uint
  for i := 0; i < len(word); i++ {
    result = 31 * result + uint(word[i])
  }
  return result
}


func htableInsertWord(h *htable, word string) {
  position := htableWordToInt(word) % uint(h.capacity)
  if h.keys[position].containerType == "" {
    containerAdd(h.keys[position], word)
    h.frequencies[position] = 1
    h.numKeys++
    return
  } else {
    containerAdd(h.keys[position], word)
    h.frequencies[position]++
  }
}


func htableSearch(h *htable, word string) bool {
  position := htableWordToInt(word) % uint(h.capacity)

  if h.keys[position].containerType != "" {
    return containerSearch(h.keys[position], word)
  }
  return false
}


func htablePrint(h *htable) {
  for i := 0; i < h.capacity; i++ {
    if h.frequencies[i] != 0 {
      containerPrint(h.keys[i])
    }
    fmt.Println()
  }
}
