// Source: https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_profit/

// you can use includes, for example:
#include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    // write your code in C++14 (g++ 6.2.0)
    
    if(A.empty()) return NULL;
    
    int min_value=A[0], max_profit=0;
    for(int i=0;i<A.size();i++){
        min_value = min(A[i], min_value);
        max_profit = max(max_profit, A[i]-min_value);
    }
    return max_profit;
}