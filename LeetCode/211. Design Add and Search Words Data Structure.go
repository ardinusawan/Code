// https://leetcode.com/problems/design-add-and-search-words-data-structure/

// Implement using Trie data structure

type WordDictionary struct {
    children map[string]*WordDictionary
    endOfWord bool
}


func Constructor() WordDictionary {
    return WordDictionary{}
}


func (this *WordDictionary) AddWord(word string)  {
    current := this
    
    for _, v := range word {
        c := string(v)
        if _, ok := current.children[c]; !ok {
            if len(current.children) == 0 {
                current.children = make(map[string]*WordDictionary)
            }
            current.children[c] = &WordDictionary{}
        }
        current = current.children[c]
    }
    current.endOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
    return searchDfs(this, word, 0)
}

func searchDfs(root *WordDictionary, word string, index int) bool {
    current := root
    
    for i:=index;i<len(word);i++ {
        c := string(word[i])
        
        // if it ".", find next match child recursive
        if c == "." {
            for _, child := range current.children {
                if searchDfs(child, word, i+1) {
                    return true
                }
            }
            return false
        } else {
            // base case
            if _, ok := current.children[c]; !ok {
                return false
            }
            current = current.children[c]
        }
    }  
    return current.endOfWord
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
