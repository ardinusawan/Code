https://www.hackerrank.com/challenges/array-left-rotation/submissions/code/34911417
#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    int a,b;
    cin >> a >> b;
    int c[a];
    for(int i=0;i<a;i++){
        cin >> c[i];
    }
    for(int j=b;j<a+b;j++){
        cout << c[j%a] << ' ';
    }
    return 0;
}

