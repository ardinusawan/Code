

// you can use includes, for example:
// #include <algorithm>
#include <set>
// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    // write your code in C++14 (g++ 6.2.0)
    set<int> s;
    for(auto& n:A){
        s.insert(n);
    }
    vector<int> v;
    for(auto& n:s){
        if(n>0) {v.push_back(n);}
    }
    if(v.size()==0||v[0]!=1) return 1;
    else if(v.size()==1&&v[0]==1) {return 2;}

    int res = 1;
    if(v.size()>1){
        for(int i=0;i<v.size();i++){
            if(v[i]>0&&v[i+1]!=v[i]+1){
                res = v[i]+1; break;
            }
        }
    }
    return res;
    
}

