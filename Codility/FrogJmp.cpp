

// you can use includes, for example:
// #include <algorithm>
#include <math.h>
// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(int X, int Y, int D) {
    double jump;
    jump = Y-X;
    jump = ceil(jump / D);
    //cout << jump;
    return jump;
}

