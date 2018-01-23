https://www.hackerearth.com/practice/algorithms/graphs/graph-representation/tutorial/
#include <iostream>
#include <vector>
using namespace std;

int main()
{
	int  n,m;
	cin >> n >> m;
	bool v[n][m];
	for(int i=0;i<n;i++){
        for(int j=0;j<m;j++){
            v[i][j]=false;
        }
	}
	for(int i=0;i<m;i++){
		int  a,b;
		cin >> a >> b;
		v[a][b]=true;
		v[b][a]=true;
	}
	int t;
	cin >> t;
	for(int i=0;i<t;i++){
        int a,b;
        cin >> a >> b;
        if(v[a][b]==true||v[b][a]==true){
            cout << "YES"<< endl;
        }
        else {
            cout << "NO" << endl;
        }
	}

    return 0;
}

