// Source : https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/count_factors/

// you can use includes, for example:
// #include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;
#include <math.h>
int solution(int N) {
    // write your code in C++14 (g++ 6.2.0)
    if(N<1) return 0;
    
    else if(N==1) return 1;
    else {
        int count = 0;
        int data = floor(sqrt(N));
        for(int i=1;i<=data;i++){
            if(i*i<N&&N%i==0){
                count+=2;    
            }
            else if(i*i==N) count+=1;

        }
        return count;
    }
}
