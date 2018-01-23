// you can use includes, for example:
#include <algorithm>
#include <set>
// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    // write your code in C++14 (g++ 6.2.0)
    set<int> s;
    for(auto& n:A){
        s.insert(n);    
    }
    return s.size(); 
}