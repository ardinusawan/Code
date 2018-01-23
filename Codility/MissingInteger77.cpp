// you can use includes, for example:
// #include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;

int solution(vector<int> &A) {
    //cari nilai terbesar positif
    //bikin array sejumlah nilai terbesar positif
    //init semua dengan 0, dengan loop saja
    //loop A, data = index. tandai index dengan 1 jika terdapat ada
    int max = 0;
    for(auto& n:A){
        if(n>0&&max<n) max = n;
        }
        
    if(max==0) return 1;
    long long data[max+2] = {0};
    for(int i=0;i<=max+1;i++){
        data[i]=0;
    }
    for(auto& n:A){
        //cout << n;
        data[n]=1;    
    }
    
    for(int i=1;i<=max+1;i++){
        if(data[i]==0) return i;
        //cout << data[i];
    }
    
}