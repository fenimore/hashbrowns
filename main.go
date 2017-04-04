// package hashbrown is a simple implementation of
// a distributed hash table.
// Nodes are "servers", each with a hash table
// So all the nodes are connected as a circular linked list
// The keysize is 128
package main

import (
	"fmt"
)

var keysize = 128

// TODO: should maxsize be 2**128?
var maxsize = keysize

// if keysize is 128...

type Node struct {
	id   int64
	data map[int64]string
	next *Node
}

// NearNode finds the nearest node to key
func NearNode(prev *Node, next *Node, key int64) *Node {
	if prev.id > next.id {
		// reached last node
	} else {
		// check which node is closer
		if key-prev.id < next.id-key {
			return prev
		} else {
			return next
		}
	}
	return prev
}

// FindNode finds the relevant node for key
func FindNode(n *Node, key int64) *Node {
	now := n
	for !(key >= now.id && key < now.next.id) {
		if key > now.next.id && key > now.id {
			// reach end of linked list
			break
		}
		now = now.next
	}
	return NearNode(now, now.next, key)
}

func Set(n *Node, key int64, val string) {
	n = FindNode(n, key)
	n.data[key] = val
}

func Get(n *Node, key int64) string {
	node := FindNode(n, key)
	return node.data[key]
}

// TODO: implement Del()

func main() {
	fmt.Println("HashBrowns!")
}
