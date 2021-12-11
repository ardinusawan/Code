// https://leetcode.com/problems/course-schedule/

var (
    cp map[int][]int
    visited map[int]bool
)

func canFinish(numCourses int, prerequisites [][]int) bool {
    // build hash map for course and prerequisites
    cp = make(map[int][]int)
    
    // build hashmap for visited node
    visited = make(map[int]bool)
    // if visited twice, return false (loop)
    
    // use recursive to check if prerequisites can be done
    // can be done = not have prerequisites
    for _, p := range prerequisites {
        cp[p[0]] = append(cp[p[0]], p[1])
    }
    
    for k, v := range cp {
        for _, v1 := range v {
            if ok := checkDfs(k, []int{v1}); !ok {
                return false
            }
        }
    }
    
    return true
}

func checkDfs(c int, pr []int) bool {
    if _, ok := visited[c]; ok {
        return false
    }
    visited[c] = true
    
    // base
    if len(pr) == 0 {
        visited = make(map[int]bool)
        return true
    }
    
    for i, _ := range pr {
        return checkDfs(pr[i], cp[pr[i]])
    }
    return true
}
