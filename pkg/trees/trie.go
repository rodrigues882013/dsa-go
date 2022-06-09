package trees

type TrieOperation interface {
	Add(word string)
	Search(word string) bool
	SearchByPrefix(word string) bool
}

type Trie struct {
	Root *TrieNode
}

func (t *Trie) Add(word string)  {
	current := t.Root

	for _, letter := range word {
		if node, ok := current.Children[letter]; ok {
			//do something here
			current = node
		} else {
			node = &TrieNode{
				Children: make(map[rune]*TrieNode),
			}
			current.Children[letter] = node
		}
	}

	current.IsEnd = true
}

func (t *Trie) Search(word string) bool {
	panic(any("Implement me"))
}

func (t *Trie) SearchByPrefix(word string) bool {
	panic(any("Implement me"))
}


