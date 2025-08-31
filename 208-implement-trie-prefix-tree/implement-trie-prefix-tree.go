type Trie struct {
	child [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{
		child: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	current := this
	for i := range word {
		if current.child[word[i]-'a'] == nil {
			newNode := Constructor()
			current.child[word[i]-'a'] = &newNode
		}
		current = current.child[word[i]-'a']
	}
	current.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for i := range word {
		if curr.child[word[i]-'a'] == nil {
			return false
		}
		curr = curr.child[word[i]-'a']
	}
	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for i := range prefix {
		if curr.child[prefix[i]-'a'] == nil {
			return false
		}
		curr = curr.child[prefix[i]-'a']
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */