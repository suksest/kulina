#include <bits/stdc++.h>

using namespace std;

int countingValleys(int n, string s) {
    int level = 0;
    int prev = 0;
    int valleys = 0;
    for(int i = 0; i < n; i++){
        if(s[i]=='U'){
            prev = level;
            level++;
        } else {
            prev = level;
            level--;
        }
        if(i>0){
            if(level==0 && prev==-1){
                valleys++;
            }
        }
    }

    return valleys;
}

int main(){
    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    string s;
    getline(cin, s);

    int result = countingValleys(n, s);

    cout << result << "\n";

    return 0;
}