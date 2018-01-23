// you can use includes, for example:
#include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    //vector<int> A = {-10,-2,-4};
    // write your code in C++14 (g++ 6.2.0)
    sort(A.begin(), A.end());
    long long  nilai=A[0]*A[1]*A[2];
    for(int i=0;i<A.size()-2;i++){
        long long  temp = A[i]*A[i+1]*A[i+2];
        if(nilai<temp) nilai = temp;
        else if(nilai<A[0]*A[1]*A[A.size()-1]) nilai = A[0]*A[1]*A[A.size()-1];
    }
    return nilai;
}