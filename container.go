/*
* Implementation of a Wrapper class container which holds a RED_BLACK_TREE
* or a Slice
* 14/5/18
* Author Callan Taylor
*/
package main

import "fmt"


type container struct {
  containerType string
  contentsArray []string
  contentsRbt *rbt
}


func newContainer(t string) *container {
  c := container{containerType:"", contentsArray:nil, contentsRbt:nil}
  if t == "RED_BLACK_TREE" {
    c = container{containerType:t, contentsArray:nil, contentsRbt:newRbt()}
  } else {
    c = container{containerType:t, contentsArray:make([]string, 0), contentsRbt:nil}
  }
  return &c
}


func containerAdd(c *container, word string) {
  if c.containerType == "RED_BLACK_TREE" {
    c.contentsRbt = rootFix(insertWord(c.contentsRbt, word))
  } else {
    c.contentsArray = append(c.contentsArray, word)
  }
}


func containerPrint(c *container) {
  if c.containerType == "RED_BLACK_TREE" {
    preorderTraverse(c.contentsRbt, printRbt)
  } else {
    for i := 0; i < len(c.contentsArray); i++ {
      fmt.Print(c.contentsArray[i], " ")
    }
  }
}


func containerSearch(c *container, word string) bool {
  if c.containerType == "RED_BLACK_TREE" {
    return rbtSearch(c.contentsRbt, word)
  } else {
    for _, value := range c.contentsArray {
      if value == word {
        return true
      }
    }
  }
  return false
}
