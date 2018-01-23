
// you can use includes, for example:
// #include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    // write your code in C++14 (g++ 6.2.0)
    long long q=0;
    long long total=0;
    for(int i=A.size()-1;i>=0;i--){
        if(A[i]==1) q++;
        else{total+=q;}
    }
    if(total>1000000000){return -1;}
    else return total;
    
}
