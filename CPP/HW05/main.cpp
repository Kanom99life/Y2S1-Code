#include <iostream>
#include <stack>
#include <queue>
#include <string>
using namespace std;

bool isBalanced(string s) {
  stack<char> tmpStack;
  for(int i = 0; i < s.length(); i++){
    if(tmpStack.empty()){
        tmpStack.push(s[i]);
    }else if((tmpStack.top() == '(' && s[i] == ')') || (tmpStack.top() == '[' && s[i] == ']') || (tmpStack.top() == '{' && s[i] == '}')){
      tmpStack.pop();
    }else{
      tmpStack.push(s[i]);
    }
  }
  if(tmpStack.empty()){
    return true;
  }
  return false;
}

int main() {
  string s;
  cin >> s;
  while (s != "-1") {
    if(isBalanced(s)) {
      cout << "Parentheses are balanced" << endl;
    } else {
      cout << "Parentheses are not balanced" << endl;
    }
    cin >> s;
  }
}