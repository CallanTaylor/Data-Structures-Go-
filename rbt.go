/*
 * Implementation of a RED_BLACK_TREE
 * 14/05/18
 *  Author Callan Taylor
 */
package main

import "fmt"
import "strings"

type fn func(*rbt)


type rbt struct {
  key string
  left *rbt
  right *rbt
  colour string
}


func isRed(node *rbt) bool {
  if node != nil && node.colour == "RED" {
    return true
  } else {
    return false
  }
}


func isBlack(node *rbt) bool {
  if node == nil || node.colour == "BLACK" {
    return true
  } else {
    return false
  }
}


func newRbt() *rbt {
  newRbt := rbt{key:"nil", left:newNilLeaf(), right:newNilLeaf(), colour:"BLACK"}
  return &newRbt
}


func newNilLeaf() *rbt {
  node := rbt{key:"nil", left:nil, right:nil, colour:"BLACK"}
  return &node
}


func rightRotate(node *rbt) *rbt {
  temp := node
  node = node.left
  temp.left = node.right
  node.right = temp
  return node
}


func leftRotate(node *rbt) *rbt {
  temp := node
  node = node.right
  temp.right = node.left
  node.left = temp
  return node
}


func rootFix(root *rbt) *rbt {
  root.colour = "BLACK"
  return root
}


func rbtFix(node *rbt) *rbt {
    if isRed(node.left) && isRed(node.left.left) {
      if isRed(node.right) {
        node.colour = "RED"
        node.left.colour = "BLACK"
        node.right.colour = "BLACK"
      } else if isBlack(node.right) {
        node = rightRotate(node)
        node.colour = "BLACK"
        node.right.colour = "RED"
      }
    } else if isRed(node.left) && isRed(node.left.right) {
      if isRed(node.right) {
        node.colour = "RED"
        node.left.colour = "BLACK"
        node.right.colour = "BLACK"
      } else if isBlack(node.right) {
        node.left = leftRotate(node.left)
        node = rightRotate(node)
        node.colour = "BLACK"
        node.right.colour = "RED"
      }
    } else if isRed(node.right) && isRed(node.right.left) {
      if isRed(node.left) {
        node.colour = "RED"
        node.left.colour = "BLACK"
        node.right.colour = "BLACK"
      } else if isBlack(node.left) {
        node.right = rightRotate(node.right)
        node = leftRotate(node)
        node.colour = "BLACK"
        node.left.colour = "RED"
      }
    } else if isRed(node.right) && isRed(node.right.right) {
      if isRed(node.left) {
        node.colour = "RED"
        node.left.colour = "BLACK"
        node.right.colour = "BLACK"
      } else if isBlack(node.left){
        node = leftRotate(node)
        node.colour = "BLACK"
        node.left.colour = "RED"
      }
    }
    return node
}


func insertWord(thisTree *rbt, word string) *rbt {
  if thisTree.key == "nil" {
    thisTree := rbt{key:word, left:newNilLeaf(), right:newNilLeaf(), colour:"RED"}
    return &thisTree
  }

  compare := strings.Compare(thisTree.key, word)

  if compare > 0 {
    thisTree.left = insertWord(thisTree.left, word)
  } else if compare < 0 {
    thisTree.right = insertWord(thisTree.right, word)
  } else if compare == 0 {
    thisTree.left = insertWord(thisTree.left, word)
  }
  thisTree = rbtFix(thisTree)
  return thisTree
}


func rbtSearch(node *rbt, word string) bool {
  if node.key == "nil" {
    return false
  }

  compare := strings.Compare(node.key, word)

  if compare == 0 {
    return true
  } else if compare > 0 {
    return rbtSearch(node.left, word)
  } else if compare < 0 {
    return rbtSearch(node.right, word)
  }
  return true
}


func inorderTraverse(thisTree *rbt, function fn) {
  if thisTree.left.key != "nil" {
    inorderTraverse(thisTree.left, function)
  }
  function(thisTree)
  if thisTree.right.key != "nil" {
    inorderTraverse(thisTree.right, function)
  }
}


func preorderTraverse(thisTree *rbt, function fn) {
  function(thisTree)
  if thisTree.left.key != "nil" {
    preorderTraverse(thisTree.left, function)
  }
  if thisTree.right.key != "nil" {
    preorderTraverse(thisTree.right, function)
  }
}


func printRbt(thisTree *rbt) {
  fmt.Println(thisTree.colour, ":  ", thisTree.key)
}
