// you can use includes, for example:
// #include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {

    int max = 0;
    for(auto& n:A){
        if(n>100000) return 0;
        else if(n>0&&max<n) max = n;
        }
        
    if(A.empty()||max==0) return 0;
    long long data[max] = {0};
    for(int i=0;i<=max;i++){
        data[i]=0;
    }
    for(auto& n:A){
        //cout << n;
        //if(n>100000||data[n]==1) return 0;
        if(data[n]==0) {data[n]=1;}
        else if(data[n]==1) {return 0;}
    }
    int check = 0;
    for(int i=1;i<=max;i++){
        if(data[i]==1) check++;
        //cout << data[i];
    }
    if(check==max) return 1;
    else return 0;
    
}