// Copyright 2015 Lars Wiegman. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package trie implements an ordered tree data structure optimized for key
// retrieval.

package trie

import (
	"errors"
	"fmt"
	"io"
)

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrKeyLength   = errors.New("key length cannot be zero")
)

// Node is a node of a trie tree.
type Node struct {
	label    rune
	children []*Node
	parent   *Node
	leaf     bool
}

// Trie represents an ordered tree data structure optimized for retrieval of
// strings stored in a dynamic set. The zero value for Trie is an empty trie
// ready to use.
type Trie struct {
	root Node
}

// IsLeaf returns true when node is also a leaf.
func (n *Node) IsLeaf() bool {
	return n.leaf
}

// Lookup returns true and the associated node when the key can be found in
// the trie.
func (t *Trie) Lookup(key string) (*Node, bool) {
	a := []rune(key)
	return t.root.Lookup(a)
}

// Lookup returns true and the associated node when the sequence of runes can
// be found in the node.
func (n *Node) Lookup(a []rune) (*Node, bool) {
	if len(a) < 1 {
		return n, false
	}
	for _, c := range n.children {
		if c.label == a[0] {
			if len(a) > 1 {
				return c.Lookup(a[1:])
			}
			return c, c.IsLeaf()
		}
	}
	return n, false
}

// Insert adds the given key to the trie.
func (t *Trie) Insert(key string) error {
	if len(key) < 1 {
		return ErrKeyLength
	}
	a := []rune(key)
	return t.root.Insert(a)
}

// Insert appends the given sequence of runes to the node.
func (n *Node) Insert(a []rune) error {
	for _, c := range n.children {
		if c.label == a[0] {
			if len(a) > 1 {
				return c.Insert(a[1:])
			}
			return nil
		}
	}
	newChild := &Node{label: a[0], parent: n}
	n.children = append(n.children, newChild)
	if len(a) > 1 {
		return newChild.Insert(a[1:])
	}
	newChild.leaf = true
	return nil
}

// Delete removes the given key.
func (t *Trie) Delete(key string) error {
	if len(key) < 1 {
		return ErrKeyLength
	}
	a := []rune(key)
	n, ok := t.root.Lookup(a)
	if !ok {
		return ErrKeyNotFound
	}
	n.leaf = false
	n.Delete()
	return nil
}

// Delete removes the node from its parent. Any node rendered obsolete by this
// is also removed.
func (n *Node) Delete() {
	if n.IsLeaf() {
		return
	}
	if len(n.children) > 0 {
		return
	}
	// remove child from parent
	var a []*Node
	for _, c := range n.parent.children {
		if c != n {
			a = append(a, c)
		}
	}
	n.parent.children = a
	n.parent.Delete()
}

// DumpKeys writes the keys from the given trie to the given Writer. The keys
// are seperated by the given separator string.
func DumpKeys(out io.Writer, sep string, t Trie) error {
	return t.root.DumpKeys(out, sep, nil)
}

// DumpKeys writes any leaf from the node to the given Writer. The leags are
// seperated by the given separator string.
func (n *Node) DumpKeys(out io.Writer, sep string, prefix []rune) error {
	if n.parent != nil {
		prefix = append(prefix, n.label)
	}
	if n.IsLeaf() {
		if _, err := fmt.Fprint(out, string(prefix), sep); err != nil {
			return err
		}
	}
	for _, c := range n.children {
		c.DumpKeys(out, sep, prefix)
	}
	return nil
}
