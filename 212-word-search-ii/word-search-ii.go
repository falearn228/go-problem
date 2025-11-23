func findWords(board [][]byte, words []string) []string {
    trie := NewTrie()
    for _, w := range words {
        trie.Set(w)
    }
    m, n := len(board), len(board[0])
    result := make([]string, 0 , len(words))

    var dfs func(r, c int, node *WordDictionary)
    dfs = func(r, c int, node *WordDictionary) {
        char := board[r][c]

        if char == '$' || node.childs[char-'a'] == nil {
            return
        }

        node = node.childs[char-'a']

        if node.word != "" {
            result = append(result, node.word)
            node.word = ""
        }

        board[r][c] = '$'
        dirs := []int{0, 1, 0, -1, 0} 
        for i := range 4 {
            nr, nc := r+dirs[i], c+dirs[i+1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n {
				dfs(nr, nc, node)
			}
        }

        board[r][c] = char
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
			dfs(i, j, trie)
		}
	}

	return result
}

type WordDictionary struct {
    childs [26]*WordDictionary
    word string
}


func NewTrie() *WordDictionary {
    return &WordDictionary{}
}


func (this *WordDictionary) Set(word string)  {
    curr := this

    for i := range word {
        idx := word[i]-'a'
        if curr.childs[idx] == nil {
            curr.childs[idx] = &WordDictionary{}
        }
        curr = curr.childs[idx]
    }

    curr.word = word
}
