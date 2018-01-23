//find max sum of binary 0 between 1 gap
#include <bits/stdc++.h>
using namespace std;
int  solution(int data) {
    int max = 0, now = 0;
    //unsigned data;
    //cin >> data;
    string biner = bitset<32>(data).to_string();
    for(int i=0;i<biner.size();i++){
    	if(max==0&&biner[i]=='1'){
    		i++;
    		while(true){
    			if(biner[i]=='0'){
    				now++;
    				i++;
    			}
    			else {
						if (max<now){
	    					max = now;
    					}
    					break;
    				}
    		}

    	}

    }
    return max;
}

int main(){
	int a;
	cin >> a;
	cout << solution(a);
}