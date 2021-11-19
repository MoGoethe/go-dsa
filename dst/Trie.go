package dst

// is the number of poosible characters in the tire
const alphabetSize = 26

type TrieNode struct {
	children [alphabetSize]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func InitTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}

func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &TrieNode{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

// 打印
//
