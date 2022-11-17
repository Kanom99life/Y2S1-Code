#include <bits/stdc++.h>
using namespace std;
class Team{
    public:
    int ID;
    int GD;//Goal Difference = Goal score - Goal conced
    int PTS;//Point (win = 3, draw = 1, lose = 0)

    //Constructor
    Team(int x, int y, int z) : ID(x), GD(y), PTS(z) {}
	
    /*
         WRITE YOUR CODE HERE
    */
   
};

class Scoreboard{
public:
    vector<Team> T;
    int numberOfTeams = 0;
    
    Scoreboard(int n){
        numberOfTeams = n;
        for(int i=0;i<n;i++){
            Team x(i,0,0);
            T.push_back(x);
        }
    
    }

    void match(int ID1, int ID2,int G1,int G2){
        int j,k;
        for (int i = 0; i<T.size(); i++){
            if(T[i].ID == ID1){
                k = i;
            }
            if(T[i].ID == ID2){
                j = i;
            }
        }

        int GD = 0;
        if (G1 > G2){
            GD = G1 - G2;
            T[k].PTS += 3;
            T[k].GD += GD;
            T[j].GD -= GD;
            
        }
            
        else if (G2 > G1){
            GD = G2 - G1;
            T[j].PTS += 3;
            T[j].GD += GD;
            T[k].GD -= GD;
           
        }
        else{
            T[k].PTS += 1;
            T[j].PTS += 1;
            
        }
        bubbleSort();
        
    } 
    void showTeamAtRank(int i){
        
        i--;
        cout << T[i].ID << " ";
        cout << T[i].PTS << " ";
        cout << T[i].GD << " ";
        cout<< "\n";
    }

    void bubbleSort(){
        int i, j, tempID, tempPTS, tempGD;
        for(i = (T.size() - 1);i >= 0;i--){
           for(j = 1; j <= i; j++){
                if(T[j-1].PTS == T[j].PTS){
                    if(T[j-1].GD < T[j].GD){
                        tempID = T[j-1].ID;
                        tempPTS = T[j-1].PTS;
                        tempGD = T[j-1].GD;
                        T[j-1] = T[j];
                        T[j].ID = tempID;
                        T[j].PTS = tempPTS;
                        T[j].GD = tempGD;
                    }
                    else if (T[j-1].GD == T[j].GD){
                        if(T[j-1].ID > T[j].ID){
                            tempID = T[j-1].ID;
                            tempPTS = T[j-1].PTS;
                            tempGD = T[j-1].GD;
                            T[j-1] = T[j];
                            T[j].ID = tempID;
                            T[j].PTS = tempPTS;
                            T[j].GD = tempGD;
                        }
                    }
                }
                else if(T[j-1].PTS < T[j].PTS){
                    tempID = T[j-1].ID;
                    tempPTS = T[j-1].PTS;
                    tempGD = T[j-1].GD;
                    T[j-1] = T[j];
                    T[j].ID = tempID;
                    T[j].PTS = tempPTS;
                    T[j].GD = tempGD;
                }
               
           }
              

            

        }
    }
};