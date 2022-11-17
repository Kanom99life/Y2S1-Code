#include <iostream>
#include "HW03.cpp"
using namespace std;

int main()
{
    Polynomial f1;
    Polynomial f2;
    f1.addTerm(5,8);
    f1.addTerm(4,3);
    
    f1.printPolynomial();
    f2.addTerm(4,10);
    f2.addTerm(6,9);

    
    // f2.addTerm(5,7);
   
    
    f2.printPolynomial();
    //f1.plus(f2);
    f1.minus(f2);
    f1.printPolynomial();
    return 0;
}
