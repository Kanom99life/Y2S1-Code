#include <iostream>
#include "term.cpp"
using namespace std;
class Polynomial{
public:
    Poly_node* head = new Poly_node;
    int size;
    Polynomial(){
        head = NULL;
        size = 0;
    }
    void addTerm(int coef, int exponent){
        Poly_node *poly = new Poly_node;
        Poly_node *cur = head;
        poly -> coef = coef;
        poly -> exponent = exponent;
        poly -> next = NULL;
        if(size == 0){
            poly -> next = head;
            head -> next == NULL;
            head = poly;
            size ++;
        }else{
           if (exponent > head -> exponent){
                poly -> next = head ;
                head = poly;
                size ++;
            }else if ( exponent == head -> exponent){
                head -> coef += coef;
                delete poly;
            }
            else{
                while(cur -> next != NULL){
                    if(cur -> next -> exponent > exponent){
                        cur = cur -> next;
                    }else{
                        break;
                    }                
                }
                if(cur -> next == NULL){
                    cur -> next = poly;
                    size ++;
                }
                else if (exponent != cur -> next -> exponent){
                    poly -> next = cur -> next;
                    cur -> next = poly;
                    size ++;
                }else if (exponent == cur -> next -> exponent){
                    cur -> next -> coef += coef;
                    delete poly;
                }    
            }
        }  
    }
    
    void printPolynomial(){
        cout<<"[ ";
        Poly_node* ptr = head;
        while(ptr!=NULL){
            if(ptr->coef==0){

            }else{
                cout<< ptr->coef<<"X^{"<<ptr->exponent<<"} ";
            }
            ptr = ptr->next;
        }
        cout<<"]\n";
    }
    void plus(Polynomial f2){
        Poly_node *cur = f2.head;
        while(cur != NULL){
            int p_coef = cur -> coef;
            int p_exponent = cur -> exponent;
            addTerm(p_coef,p_exponent);
            cur = cur -> next; 
        }
    }
    void minus(Polynomial f2){
       Poly_node *cur = f2.head;
        while(cur != NULL){
            int p_coef = (cur -> coef) * -1;
            int p_exponent = cur -> exponent;
            addTerm(p_coef,p_exponent);
            cur = cur -> next; 
        }
    }
};