// Copyright 2015 Lars Wiegman. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package trie implements an ordered tree data structure optimized for key
// retrieval.

package trie

import (
	"bytes"
	"testing"
)

var data = []string{
	"go",
	"goad",
	"goaded",
	"goading",
	"goads",
	"goal",
	"goaled",
	"goalie",
	"goalies",
	"goaling",
	"goalkeeper",
	"goalkeepers",
	"goalless",
	"goalpost",
	"goalposts",
	"goals",
	"goaltender",
	"goaltenders",
}

func TestInsert(t *testing.T) {
	var testTable = []struct {
		key      string
		expected bool
	}{
		{"go", true},
		{"goad", true},
		{"goat", false},
		{"oat", false},
		{"", false},
	}

	root := Trie{}
	for _, s := range data {
		if err := root.Insert(s); err != nil {
			t.Fatal(err)
		}
	}

	for _, test := range testTable {
		_, result := root.Lookup(test.key)
		if test.expected != result {
			t.Errorf("Result should have been %t, but it was %t for %s", test.expected, result, test.key)
		}
	}
}

func TestDelete(t *testing.T) {
	var testTable = []struct {
		key      string
		expected bool
	}{
		{"go", false},
		{"goal", false},
		{"goalpost", false},
		{"goalkeepers", false},
	}

	root := Trie{}
	for _, s := range data {
		if err := root.Insert(s); err != nil {
			t.Fatal(err)
		}
	}

	for _, test := range testTable {
		err := root.Delete(test.key)
		if err != nil {
			t.Fatal(err)
		}
		_, result := root.Lookup(test.key)
		if test.expected != result {
			t.Errorf("Result should have been %t, but it was %t", test, result)
		}
	}
}

func TestErr(t *testing.T) {
	var testTable = []struct {
		key      string
		expected error
	}{
		{"", ErrKeyLength},
		{"_", ErrKeyNotFound},
	}

	root := Trie{}
	for _, s := range data {
		if err := root.Insert(s); err != nil {
			t.Fatal(err)
		}
	}

	for _, test := range testTable {
		result := root.Delete(test.key)
		if test.expected != result {
			t.Errorf("Result should have been %v, but it was %v", test, result)
		}
	}
}

func TestDumpKeys(t *testing.T) {
	root := Trie{}
	for _, s := range data {
		if err := root.Insert(s); err != nil {
			t.Fatal(err)
		}
	}

	buf := bytes.NewBuffer(nil)
	if err := DumpKeys(buf, "", root); err != nil {
		t.Fatal(err)
	}
	for _, key := range data {
		s := string(buf.Next(len(key)))
		if key != s {
			t.Errorf("Result should have been %q, but it was %q", key, s)
		}
	}
	if n := len(buf.Next(1)); n != 0 {
		t.Errorf("Result should have been %q, but it was %q", 0, n)
	}
}
