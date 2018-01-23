// you can use includes, for example:
// #include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    // write your code in C++14 (g++ 6.2.0)
    
    int mid = A[0];
    int ctr = 1;
    
    for(int i=1;i<A.size();i++){
        if(A[i]==mid) ctr++;
        else ctr--;
        if(ctr==0){
            ctr = 1;
            mid = A[i];
        } 
    }
    
    int total = 0;
    for(int i=0;i<A.size();i++){
        if(A[i]==mid) total++;
    }
    
    if(total<A.size()/2) return 0;
    
    int sum = 0;
    int sum_left = 0;
    int sum_right = 0;
    for(int i=0;i<A.size();i++){
        if(A[i]==mid) sum_left ++;
        sum_right = total - sum_left;
        if((sum_left>(i+1)/2) && (sum_right>(A.size()-1-i)/2)) sum++;
    }
    return sum;
    
}

//inspired by: https://rafal.io/posts/codility-equi-leader.html