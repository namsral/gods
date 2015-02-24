Trie Data Structure
===================

Package trie implements an ordered tree data structure optimized for key retrieval.

Example:

```go
var data = []string{
	"go",
	"goad",
	"goaded",
	"goading",
	"goads",
	"goal",
	"goaled",
}

root := Trie{}
for _, s := range data {
 	root.Insert(s)
}

node, ok := root.Lookup("go")
if ok {
	fmt.Print("key go was found")
}
```

For more information about the trie data structure see the [Wikipedia article][0].

[0]: http://en.wikipedia.org/wiki/Trie "Trie"