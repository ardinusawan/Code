// https://leetcode.com/problems/implement-trie-prefix-tree/

// create two map to store word and prefix
// every time insert new data, populate those two
// time and space complexity of insert O(n)
// space and time complexity search and find prefix O(1) 

type Trie struct {
    words map[string]bool
    prefix map[string]bool    
}


func Constructor() Trie {
    return Trie{
        words: make(map[string]bool),
        prefix: make(map[string]bool),
    } 
}


func (this *Trie) Insert(word string) {
    if this.Search(word) == false {
        // populate prefix
        for i, _ := range word {
            this.prefix[word[:i+1]] = true
        }
        
        this.words[word] = true
    }
}


func (this *Trie) Search(word string) bool {
    return this.words[word]
}


func (this *Trie) StartsWith(prefix string) bool {
   return this.prefix[prefix]
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
