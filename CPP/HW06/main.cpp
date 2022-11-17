#include <bits/stdc++.h>
#include "HW06.cpp"
using namespace std;

int main(){
    int n = 4;
    //cin>>n;
    Scoreboard S(n);
    S.showTeamAtRank(1);
    S.showTeamAtRank(2);
    S.showTeamAtRank(3);
    S.showTeamAtRank(4);

    cout << "----------------------------" << endl;
    S.match(0,1,2,3);
    S.showTeamAtRank(1);
    S.showTeamAtRank(2);
    S.showTeamAtRank(3);
    S.showTeamAtRank(4);

    cout << "----------------------------" << endl;
    S.match(3,2,1,1);
    S.showTeamAtRank(1);
    S.showTeamAtRank(2);
    S.showTeamAtRank(3);
    S.showTeamAtRank(4);

    cout << "----------------------------" << endl;
    S.match(0,2,3,1);
    S.showTeamAtRank(1);
    S.showTeamAtRank(2);
    S.showTeamAtRank(3);
    S.showTeamAtRank(4);


    
    
    



    
    return 0;
}