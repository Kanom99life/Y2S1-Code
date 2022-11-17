#include <iostream>
#include <string>
#include "card.cpp"

using namespace std;

class CardList{

public:
	Card *head;
    int size = 0;
    CardList(){
         head = NULL ;
         
    }

	void pop_back(){
        
        Card *curr = head;
        Card *prev = head;
        if (size == 1){
            size--;
            head = head -> next;
            delete curr;
            curr = NULL;
        }
        else if (size > 1){
            while(prev->next->next != NULL){
                prev = prev -> next;
            }
            size--;
            curr = prev -> next;
            prev -> next = NULL;
            delete curr;
            curr = NULL;
        }
	}


    void insert_back(string newItem){
    
        Card *newCard = new Card;
        Card *curr = head;
        newCard -> name = newItem;
        newCard -> next = NULL;

        size++;
        if (size == 1){
            newCard->next = head;
            head = newCard;
        }
        else{
            while(curr -> next != NULL){
                curr = curr -> next;
            }
            curr -> next = newCard;
        }
    }
    

    void shuffle(int pos){ //1<=pos<=size-1
        Card *curr;
        Card *prev = head;
        Card *first = head;
        Card *last = head;
        if (pos >= 1 && size - 1 >= pos) {
            int i = 0;
            while (last->next != NULL) {
                last = last->next; 
                if (size- 1 == pos) {
                    if (i < size - 3) {
                        prev = prev->next;
                        i++;
                    }
                }
                else if (i < pos-1) {
                    prev = prev->next;
                }
                i++;
            }
            if (size== 2) {
                head = last;
                head->next = first;
                head->next->next = NULL;
            }
            else if (size- 1 == pos) {
                head = last;
                curr = last;
                curr->next = first;
                if (pos >= 4) { 
                    prev = prev->next;
                    for ( i = 0; i < pos-4; i++); {
                        prev = prev->next;
                    }
                    prev->next = NULL;
                }
                else {
                    prev->next->next = NULL;
                }
            }
            else {
                curr = prev->next;
                head = curr;
                prev->next = NULL;
                last->next = first;
            }
        }
    }
    
    void printCardList(){
        /*
         DO NOT DELETE OR EDIT
         */
        cout << "[ ";
        Card* ptr = head;
        while(ptr!=NULL){
            cout << ptr->name << " ";
            ptr = ptr->next;
        }
        cout << "]\n";
    }
};
