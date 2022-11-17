#include <iostream>
#include <vector>
#include <queue>
#include <algorithm>

using namespace std;
class Village{
     int numHouse;
     vector<int> *adjLists;
     bool *bfs;
     bool *dfs;

public:
     Village(int V){
          numHouse = V;
          adjLists = new vector<int>[V];
     }
     
     void addRoad(int src, int dest){
          adjLists[src].push_back(dest);
          adjLists[dest].push_back(src);
     }

     void sort_adjLists(){
          for(int i = 0; i < numHouse; i++){
               sort(adjLists[i].begin(),adjLists[i].end());
          }
     }
    
     void DFS(int firstHouse){
          dfs = new bool[numHouse];
          for (int i = 0; i < numHouse; i++){
               dfs[i] = false;
          }
      
          DFSVISIT(firstHouse);
      
          for(int i=0;i<numHouse;i++){
               if(dfs[i]==false){
               DFSVISIT(i);
               }
          }  
     }     
   
     void BFS(int firstHouse){
          bfs = new bool[numHouse];
          for (int i = 0; i < numHouse; i++){
               bfs[i] = false;
          }
          queue<int> q;
          q.push(firstHouse);
          while(!q.empty()){
               int u = q.front();
               q.pop();
               if(bfs[u]){
                    continue;
               }
               bfs[u] = true;
               cout<<u<<" ";
               sort_adjLists();
               vector<int>::iterator it;
               for(it = adjLists[u].begin(); it!=adjLists[u].end(); it++){
                    if(!bfs[*it])
                    q.push(*it);
               }
          }
     }

     void DFSVISIT(int vertex){
          dfs[vertex] = true;
          cout<< vertex<< " ";
          sort_adjLists();
          vector<int>::iterator it;
          for(it = adjLists[vertex].begin(); it!=adjLists[vertex].end(); it++){
               if(!dfs[*it])
               DFSVISIT(*it);
          }
     }
 
     void printGraph(){
          for(int i=0; i<numHouse; i++){
          vector<int> adjList = adjLists[i];
          vector<int>::iterator it;
          cout<<i<<"->";
          for(it = adjList.begin(); it!=adjList.end(); it++){
               cout<<*it<<" ";
          }
          cout<<"\n";
          }
     }
};