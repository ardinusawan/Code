// Source : https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/min_perimeter_rectangle/start/

// you can use includes, for example:
// #include <algorithm>
#include <math.h>
// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(int N) {
    // write your code in C++14 (g++ 6.2.0)
    if (N<1) return 0;
    int result = 2*(N+1);
    for(int i=1;i*i<=N;i++){
        int num1 = i;
        int num2 = N/i;
        int area = 2*(num1+num2);
        if(N%i==0&&num1*num2==N&&area<result){
            result = area;            
        }
    }
    return result;
    
    
}