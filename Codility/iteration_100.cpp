#include <bitset>
#include <vector>
using namespace std;

int solution(int N) {
    string s = bitset<64>(N).to_string();
    vector<int> v;
    int h = 0;
    for(int i=s.size()-1;i>=0;i--){
        if(s[i] == '1') {
            v.push_back(i);
            h ++;
        }
    }
    int hasil = 0;
    for(int o=0;o<h-1;o++){
        if(hasil < (v[o] - v[o+1]) -1) hasil = (v[o] - v[o+1] -1);
    }
    return hasil;
}