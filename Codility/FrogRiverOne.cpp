// you can use includes, for example:
// #include <algorithm>
#include <set>
// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(int X, vector<int> &A) {
    // write your code in C++14 (g++ 6.2.0)
    set<int> s;
    for(int i=0;i<A.size();i++){
        s.insert(A[i]);
        //cout << s.size();
        if(s.size()==X){ return i; break;} 
    }
    return -1;
    
}