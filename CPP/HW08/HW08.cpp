#include "BSTNode.cpp"
#include<iostream>
#include <bits/stdc++.h>

using namespace std;
class BST {

public:
  BSTNode* root;
  
  BST() {
    root = nullptr;
  }
  void insert(int value) {
    BSTNode *new_node = new BSTNode();
    BSTNode *p,*prev;
    new_node -> value = value;
    new_node -> left = NULL;
    new_node -> right = NULL;
    if(root == NULL){
      root = new_node;
      root -> left = NULL;
      root -> right = NULL;
    }
    else {
      p = root;
      while(p != NULL){
        prev = p;
        if(p -> value > new_node -> value){
          p = p -> left;
        }
        else {
          p = p -> right;
        }
      }
      if(prev -> value > new_node -> value){
        prev -> left = new_node;
      }
      else {
        prev -> right = new_node;
      }
    }
  }

  void remove(int value){
      delete_element(root, root, value); 
  }

  BSTNode *MinNode(BSTNode *a){
    /* loop down to find the leftmost leaf */
    while (a != nullptr && a->left != nullptr)
      a = a->left;
      return a;
  }

  void delete_element(BSTNode *a, BSTNode *before, int value) {
    if (a != nullptr) {
      if (value < a->value){
        before = a;
        delete_element(a->left, before, value);
      } 
      else if (value > a->value){
        before = a;
        delete_element(a->right, before, value);
      } 
      else{
        if (a->left == nullptr && a->right == nullptr){
          if (before->value == a->value){
            root = nullptr;
          } 
          else if (before->left->value == a->value){
            before->left = nullptr;
          }
          else{
            before->right = nullptr;
          }
        }
        else if (a->left == nullptr){
          if (before->value == a->value){
            root = a->right;
          }
          else if (before->left->value == a->value){
            before->left = a->right;     
          }
          else{
              before->right = a->right;
          }
        }
        else if (a->right == nullptr){
          if (before->value == a->value){
            root = a->left;
          } 
          else if (before->left->value == a->value){
            before->left = a->left;
          } 
          else{
            before->right = a->left;
          }
        }
        else{
          
          BSTNode *Curr = new BSTNode;
          Curr = MinNode(a->right);
          int num = Curr->value;
          delete_element(root, root, Curr->value);
          if(before->value == a->value){
            before->value = num;
          }
          else if (before->left->value == a->value){
            before->left->value = num;
          }
          else{
            before->right->value = num;
          }
        }
      }
    }
  }
  
  int get_depth(int value) {
    BSTNode *p,*prev;
    int depth = 0;
    if (value == root -> value){
      return depth;
    }
    if(!findValue(root,value)){
      return -1;
    }
    p = root;
    while (p -> value != value){
      if(p -> value > value){
        p = p -> left;
        depth++;
      }
      else {
        p = p -> right;
        depth++;
      }
    }
    return depth;
  }
  
  bool findValue(BSTNode* node, int value){
  
    if (node == NULL) 
      return false;
    if (node -> value == value) 
      return true;
    if (findValue(node->left, value)){
      return true;
    }
    if (findValue(node->right, value)){
      return true;
    }
    return false;
    
  }

  // void searchKey(BSTNode* &curr, int key, BSTNode* &parent){
    
  //   while (curr != nullptr && curr->value != key)
  //   {
        
  //       parent = curr;
 
  //       if (key < curr->value) {
  //           curr = curr->left;
  //       }
  //       else {
  //           curr = curr->right;
  //       }
  //   }
  // }

  // BSTNode* getMinimumKey(BSTNode* curr){
  //   while (curr->left != nullptr) {
  //       curr = curr->left;
  //   }
  //   return curr;
  // }
 

};