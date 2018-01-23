// you can use includes, for example:
// #include <algorithm>
#include <algorithm>
// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    // write your code in C++14 (g++ 6.2.0)
    vector<int> B = A;
    sort(A.begin(), A.end());
    int mid = A[A.size()/2];
    int size = 0;
    for(auto& n:A){
        if(n==mid) size++;
    }
    if(size>A.size()/2) {
            for(int i=0;i<A.size();i++){
                //cout << i;
                if(B[i]==mid) {
                    //cout << i << endl;
                    return i;
                }
            }
        }
    else return -1;
}