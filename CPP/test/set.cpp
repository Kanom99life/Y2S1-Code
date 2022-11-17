#include <iostream>
#include <set>
using namespace std;

class set_set{
    public:    
    set_set(){
        set<int> result;
    }

    set<int> getUnion(set<int> a, set<int> b){
        set<int> result = a;
        set<int>::iterator itb;
        for (itb = b.begin(); itb!= b.end(); itb++){
            result.insert(*itb);
        }
        return result;
    }

    void printSet(set<int> result){
        set<int>:: iterator it;
        for( it = result.begin(); it != result.end(); it++){
            cout<<*it<<" ";
        }

    }

};