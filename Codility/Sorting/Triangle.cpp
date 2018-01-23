// you can use includes, for example:
#include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    // write your code in C++14 (g++ 6.2.0)
    vector<pair<int, int> >v;
    for(int i=0;i<A.size();i++){
        v.push_back(make_pair(A[i], i));
    }
    sort(v.begin(),v.end());
    if(A.size()>=3){
        for(int i=0;i<v.size()-2;i++){
            long long p = v[i].first;
            long long q = v[i+1].first;
            long long r = v[i+2].first;
            if(p+q>r&&q+r>p&&r+p>q) return 1;
        }
    }
    return 0;
}