// you can use includes, for example:
// #include <algorithm>

// you can write to stdout for debugging purposes, e.g.
// cout << "this is a debug message" << endl;
int top = -1;

void push(int stack[], int x, int n){
    if(top==n-1){
        cout<< "Stack full";
    }
    else{
        top++;
        stack[top] = x;
    }   
}

bool isEmpty(){
    if(top == -1){
    return true;    
    }
    else return false;    
}
void pop(){
    top--;
}
int size(){
    return top+1;    
}

int topElement(int stack[]){
    return stack[top];
}
int solution(string &S) {
    // write your code in C++14 (g++ 6.2.0)
    int s[1000000];
    for(auto& n:S){
        if(n=='(' && topElement(s)!=')') {
           push(s,n,end(s)-begin(s)); 
        }
        else{
            if(isEmpty()){
                return false;
                break;
            }
            else pop();
        }
    }
    if(top==-1){
        return true;    
    }
    else return false;
}