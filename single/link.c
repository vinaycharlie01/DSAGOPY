#include <stdio.h>
#include <stdlib.h>


typedef struct node node;
struct node
{
    int data;
    node *next;   
};
node *head,*newNode,*tail ,*p,*temp;
void  append(int data){
    newNode= (node*)  malloc(sizeof(node));
    newNode->data=data;
    newNode->next = NULL;
   if (head==NULL){
    head= newNode ;
   }else{
    tail->next = newNode;
   }
   tail=newNode;
}

void display(){
    p=head;
    printf("%d\n",p->data);
    while (p!=0)
    {
       printf("%d ",p->data);
       p=p->next;
    }
    printf("\n");
    printf("%d",tail->data);
}
void insertstart(int data){
    newNode= (node*)  malloc(sizeof(node));
    newNode->data=data;
    newNode->next=head;
    head=newNode;
}
void insertend(int data){
     newNode= (node*) malloc(sizeof(node));
     newNode->data=data;
     newNode->next= NULL;
    p=head;
    while (p->next!=NULL){
       p=p->next;
    }
    p->next=newNode;
    tail=newNode;
}
void insertatpos(int data, int pos){
    newNode= (node*) malloc(sizeof(node));
    newNode->data=data;
    int i=0;
    p=head;
    while(i<pos-1&&p->next!=NULL){
        p=p->next;
        i+=1;
    }
    newNode->next=p->next;
    p->next=newNode;
}


void delectfrombig(){
    temp=head;
    head=temp->next;
    free(temp);
}
void delectfromend(){
   p=head;
   while (p->next->next!=NULL)
   {
    p=p->next;
   }
   tail=p;
   tail->next=NULL;
   free(tail->next);
   
    
}
void delectatany(int pos){
   p=head;int i=0;
   while (i<pos-1){
    p=p->next;
    i++;
   }
   printf("%d",p->data);
   temp=p->next;
   p->next=temp->next;
   free(temp);
    
}
int main(){
    int a=5;
append(10);
append(20);
append(30);
append(40);
insertstart(50);
insertstart(30);
insertstart(110);
insertstart(40);
insertend(80);
insertatpos(60,8);
delectfrombig();
delectfromend();
delectatany(3);
display();
printf("Hello, world!");
printf("Hello, world!");
    
}