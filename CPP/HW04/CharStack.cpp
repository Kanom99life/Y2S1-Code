class CharStack {

private:
  char itemStack[10];
  char item;
  int index;

public:
  
  CharStack() { // constructor
    index = -1;
  }

  void push(char new_item) {
    index++;
    itemStack[index] = new_item;
  }

  char pop() {
    if (isEmpty()){
      return false;
    }
    else{
      item = itemStack[index];
      index--; 
      return item;

    }
  }

  char top() {
    if (isEmpty()){
      return false;
    }
    else{
      item = itemStack[index];
      return item;

    }
  }

  bool isEmpty() {
    if(index == -1){
      return true;
    }
    else{
      return false;
    }
  }
};
