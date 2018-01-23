// you can use includes, for example:
// #include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(int A, int B, int K) {
    int b = B/K;
    int a = (A > 0 ? (A-1)/K:A);
    if(A==0){
        b++;
    }
    return b - a;
    
}