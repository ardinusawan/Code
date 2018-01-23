https://www.hackerrank.com/challenges/caesar-cipher-1
#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    int t;
    cin >> t;
    char a[t];
    for(int i=0;i<t;i++){
        cin >> a[i];
    }
    int add;
    cin >> add;
    for(int j=0;j<t;j++){
        if(a[j]>='A'&&a[j]<='Z'){
            int temp = ((a[j] - 0) + add%26)%91;
            if(temp<65){
                temp += 65;
            }
            cout << (char)temp;
        }
        else if(a[j]>='a'&&a[j]<='z'){
            int temp = ((a[j] - 0) + add%26)%123;
            if(temp<97){
                temp += 97;
            }
            cout << (char)temp;
        }
        else cout << a[j];
    }
    
    return 0;
}

