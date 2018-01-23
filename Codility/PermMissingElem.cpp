// you can use includes, for example:
// #include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    // write your code in C++14 (g++ 6.2.0)
    if(A.size()>0){
        int t;
        int x1=A[0];
        int x2=1;
        for(t=1;t<A.size();t++){
            x1 = x1^A[t];
        }
        for(t=2;t<=A.size()+1;t++){
            x2 = x2^t;    
        }
        return x1^x2;
    }
    else if(A.empty()) return 1;
    else return 0;
}

