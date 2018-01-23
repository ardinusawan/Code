// Source: https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_double_slice_sum/

// you can use includes, for example:
#include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    if(A.empty()) return NULL;
    
    int max1[A.size()]={0},max2[A.size()]={0};
    
    for(int i=1;i<A.size()-1;i++){
        max1[i] = max(0, max1[i-1] + A[i]);
    }
    
    for(int i=A.size()-2;i>0;i--){
        max2[i] = max(0 , max2[i+1] + A[i]);
    }
    int maxx=0;
    for(int i=1;i<A.size()-1;i++){
        maxx=max(maxx, max1[i-1]+max2[i+1]);
    }
    return maxx;
}