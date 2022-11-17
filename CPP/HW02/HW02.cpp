#include <iostream>
#include <string>
#include "station.cpp"
using namespace std;

class Trip {

public:
    Station* header = new Station;
    Station* trailer = new Station;
    int size = 0;
    Trip()
    {
         header -> prev = NULL;
         header -> next = trailer;
         trailer -> prev = header;
         trailer -> next = NULL;
    }

     bool isEmpty(){
        return(header->next==trailer);
     }

    void printList()
    {
        cout << "[ ";
        Station* curr = header->next;
        while (curr != trailer) {
            cout << curr->name << " ";
            curr = curr->next;
        }
        cout << "]\n";
    }

    void insert_front(string newItem)
    {
         size ++;
         Station *newStation = new Station;
         newStation -> name = newItem;

         if (isEmpty()){
               newStation -> next = trailer;
               newStation -> prev = header;
               header -> next = newStation;
               trailer -> prev = newStation;
         }
         else{
               newStation -> prev = header;
               newStation -> next = header -> next;
               header -> next = newStation;
               newStation -> next -> prev = newStation;
          }
     }

    void insert_back(string newItem)
    {
          size ++;
          Station *newStation = new Station;
          newStation -> name = newItem;

          if (isEmpty()){
               newStation -> next = trailer;
               newStation -> prev = header;
               header -> next = newStation;
               trailer -> prev = newStation;
         }
          else{
               newStation -> prev = trailer -> prev;
               newStation -> next = trailer;
               trailer -> prev -> next = newStation;
               trailer -> prev = newStation;         
          }
    }

    void remove_front()
    {
          if(!isEmpty()){
               Station *front;
               if( size == 1 ){
                    header -> next = trailer;
                    trailer -> prev = header;
               }
               else{
                    Station *newFront;
                    front = header -> next;
                    newFront = header -> next -> next;
                    newFront -> prev = header;
                    header -> next = newFront;
               }
               size --;
               delete front;
          }
    }
    void remove_back()
    {
         if(!isEmpty()){
               Station *back;
               if( size == 1){
                    header -> next = trailer;
                    trailer -> prev = header;
               }
               else{
                    Station *newBack;
                    back = trailer -> prev;
                    newBack = trailer -> prev -> prev;
                    newBack -> next = trailer;
                    trailer -> prev = newBack;
               }
               size --;
               delete back;
         }
    }

    Station* visit(int nStep, string stepText)
    {
          if (!isEmpty()){
               Station *curr = header->next;
               Station *last = trailer->prev;
               for (int i = 0; i < nStep; i++){
                    if (stepText[i] == 'R'){
                         if (curr->next != trailer){
                         curr = curr->next;
                         }
                    }
                    else{
                         if (curr->prev != header){
                         curr = curr->prev;
                         }  
                    }
               }        
               return curr;             
          }  
     }
};