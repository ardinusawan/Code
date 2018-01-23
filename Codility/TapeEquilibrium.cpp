// you can use includes, for example:
// #include <algorithm>
#include <stdlib.h>
// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    // write your code in C++14 (g++ 6.2.0)
    int sum_before = A[0];
    int total = 0;
    int diff = 2001;
    for (auto& n : A){
        total+= n;
    }
    for(int i=1;i<A.size();i++){
        int tmp = abs(sum_before - (total-sum_before));
        if(diff>tmp){
            diff = tmp;
        }
        sum_before+=A[i];
    }
    return diff;
}