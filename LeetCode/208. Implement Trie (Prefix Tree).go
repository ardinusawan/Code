// https://leetcode.com/problems/implement-trie-prefix-tree/

type Trie struct {
    children map[string]*Trie
    endOfWord bool
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
    current := this
    
    for _, v := range word {
        c := string(v)
        if _, ok := current.children[c]; !ok {
            if len(current.children) == 0 {
                current.children = make(map[string]*Trie)    
            }
            current.children[c] = &Trie{}
        }
     
        current = current.children[c]
    }
    current.endOfWord = true
}


func (this *Trie) Search(word string) bool {
    current := this
    result := false

    for _, v := range word {
        c := string(v)
        if _, ok := current.children[c]; ok {
            result = current.children[c].endOfWord == true
            current = current.children[c]
        } else {
            result = false
            break
        }
    }
    return result
}


func (this *Trie) StartsWith(prefix string) bool {
    current := this
    for _, v := range prefix {
        c := string(v)
        if _, ok := current.children[c]; ok {
            current = current.children[c]
        } else {
            return false
        }
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
