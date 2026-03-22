package main

type Trie struct {
	Next  []*Trie
	IsEnd bool
}

func NewTrie() *Trie {
	return &Trie{
		Next:  make([]*Trie, 30),
		IsEnd: false,
	}
}

func (t *Trie) Insert(s string) {
	curNode := t
	for _, v := range s {
		cIdx := byte(v - 'a')
		if curNode.Next[cIdx] == nil {
			curNode.Next[cIdx] = NewTrie()
		}
		curNode = curNode.Next[cIdx]
	}
	curNode.IsEnd = true
}

type StreamChecker struct {
	sTrie   *Trie
	activeNodes []*Trie
}

func Constructor(words []string) StreamChecker {
	t := NewTrie()
	for _, v := range words {
		t.Insert(v)
	}
	return StreamChecker{
		sTrie:   t,
		activeNodes: make([]*Trie, 0),
	}
}

func (c *StreamChecker) Query(letter byte) bool {
	letter = byte(letter - 'a')
	isEnd := false

	nextNodes := make([]*Trie, 0, len(c.activeNodes)+2)

	//set starting ptr (empty string)
	c.activeNodes = append(c.activeNodes, c.sTrie)

	for _, k := range c.activeNodes {
		k = k.Next[letter]
		if k != nil {
			nextNodes = append(nextNodes, k)
			if !isEnd && k.IsEnd {
				isEnd = true
			}
		}
	}

	c.activeNodes = nextNodes

	return isEnd
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
