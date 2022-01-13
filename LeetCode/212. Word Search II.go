// https://leetcode.com/problems/word-search-ii/

// use Trie data structure
// insert all words in trie node
// iterate board, find prefix using dfs.
// if it reach endOfWord, return word as result

type Trie struct {
    endOfWord bool
    children map[string]*Trie
}

func Constructor() Trie {
    return Trie{
        false,
        make(map[string]*Trie),
    }
}

func (this *Trie) Insert(word string) {
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

var (
    res = make(map[string]bool)
    //visit = make(map[int]map[int]bool)
    visit [][]int
    ROWS, COLS = 0, 0
    b [][]byte
)

func findWords(board [][]byte, words []string) []string {
    t := Constructor()

    // insert to Trie
    for _, word := range words {
        t.Insert(word)   
    } 
    
    ROWS, COLS = len(board), len(board[0])
    b = board
    visit = make([][]int, ROWS)
    for i:=0;i<ROWS;i++ {
        visit[i] = make([]int, COLS)
    }
    res = make(map[string]bool)
        
    // iterate board and find prefix in trie
    for r, _ := range board {
        for c, _ := range board[r] {
            dfs(r, c, &t, "")
        }
    } 
    
    var r []string
    for v, _ := range res {
        r = append(r, v)
    }
    return r
}


func dfs(r int, c int, node *Trie, word string) {
    if r < 0 || c < 0 || r ==  ROWS || c == COLS || visit[r][c] != 0 {
        return
    }
    
    if _, ok := node.children[string(b[r][c])]; !ok {
        return
    }
    
    visit[r][c] += 1
    node = node.children[string(b[r][c])]
    word += string(b[r][c])
    if node.endOfWord {
        res[word] = true
    }
    
    dfs(r, c+1, node, word)    
    dfs(r+1, c, node, word)    
    dfs(r-1, c, node, word)    
    dfs(r, c-1, node, word)    
    visit[r][c] -= 1
}
