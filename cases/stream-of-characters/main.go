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
	sTrie  *Trie
	ptrSet map[*Trie]struct{}
}

func Constructor(words []string) StreamChecker {
	t := NewTrie()
	for _, v := range words {
		t.Insert(v)
	}
	return StreamChecker{
		sTrie:  t,
		ptrSet: make(map[*Trie]struct{}),
	}
}

func (this *StreamChecker) Query(letter byte) bool {
	letter = byte(letter - 'a')
	isEnd := false

	newPtrSet := make(map[*Trie]struct{})

	//set starting ptr (empty string)
	this.ptrSet[this.sTrie] = struct{}{}

	for k := range this.ptrSet {
		k = k.Next[letter]
		if k != nil {
			newPtrSet[k] = struct{}{}
			if !isEnd && k.IsEnd {
				isEnd = true
			}
		}
	}

	this.ptrSet = newPtrSet

	return isEnd
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
