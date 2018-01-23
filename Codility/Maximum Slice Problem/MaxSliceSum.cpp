// Source: https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_slice_sum/

// you can use includes, for example:
// #include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    // write your code in C++14 (g++ 6.2.0)
    if(A.empty()) return NULL;

    int max_current = 0;
    int max_global = A[0];
    for(auto &n:A){
        max_current = (n > max_current + n) ? n : max_current + n;
        max_global = (max_global > max_current) ? max_global : max_current;
    }
    
    return max_global;
}