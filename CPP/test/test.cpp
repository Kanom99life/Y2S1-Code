#include <iostream>
#include <set>
using namespace std;

template<typename T>

set<T> getUnion(set<T> a, set<T> b)
{  
    typename set<T>::iterator itb;
    for (itb = b.begin(); itb != b.end(); itb++)
    {
        a.insert(*itb);
    }
    return a;
}


template<typename T>
set<T> getInter(set<T> a, set<T> b){
    set<T> result;
    typename set<T>::iterator it;
    for ( it = a.begin() ; it != a.end(); it++){
        if(b.find(*it) != b.end()){
            result.insert(*it);
        }
    }
    return result;
}


template<typename T>
set<T> getDiff(set<T> a, set<T> b){
    set<T> result = a ;
    typename set<T>::iterator it;
    for (it = b.begin(); it != b.end(); it++){
        result.erase(*it);
    }
    return result;
}

template<typename T>
void printSet(set<T> result)
{
    typename set<T>::iterator it;
    for (it = result.begin(); it != result.end(); it++)
    {
        cout << *it << " ";
    }
    cout << endl;
}

template<typename T>
set<T> addAll(set<T>a, set<T>b){
    typename set<T> :: iterator it;
    for(it = b.begin(); it != b.end(); it++){
        a.insert(*it);
    }
    return a;
}

template<typename T>
bool containAll(set<T>a, set<T>b){
    typename set<T>::iterator it;
    for(it = b.begin(); it != b.end(); it++){
        if(a.find(*it) == a.end()){
            return false;
        }
    }
    return true;
}

template<typename T>
set<T> retainAll(set<T> a, set<T> b){
    set<T> result;
    typename set<T>:: iterator it;
    for( it = a.begin(); it != a.end(); it++){
        if(b.find(*it) != b.end()){
            result.insert(*it);
        }
    }
    return result;
}

template<typename T>
set<T>removeAll(set<T> a, set<T> b){
    set<T> result = a;
    typename set<T>::iterator it;
    for(it = b.begin(); it != b.end(); it++){
        result.erase(*it);   
    }
    return result;
}

int main()
{

    set<char> s1 = {'1','2','5'};
    set<char> s2 = {'1','2','3'};
    set<char> result = removeAll(s1, s2);
    printSet(result);
    //cout<<containAll(s1, s2)<<endl;
    return 0;
}