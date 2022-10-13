#include <stdio.h>
#include <stdlib.h>
typedef struct Node Node;
typedef struct LinkedList LinkedList;
struct Node {
    int data;
    Node* next;
};

struct LinkedList {
    Node* head;
    Node* tail;
};
Node* p;
// void free_singly_linked_list(Node* node) {
//     while (node) {
//         Node* temp = node;
//         node = node->next;
//         free(temp);
//     }
// }
Node* create(int node_data) {
    Node* node =(Node*) malloc(sizeof(Node));
    node->data = node_data;
    node->next = NULL;
    return node;
}

// void insert(LinkedList* L, int node_data) {
//     Node* node = create(node_data);
//     if (!L->head) {
//         L->head = node;
//     } else {
//         L->tail->next = node;
//     }
//     L->tail = node;
// }
void insertat(LinkedList* llist, int node_data) {
    Node* node = create(node_data);
    if (llist->head==NULL ){
        (*llist).head=node;
    }else{
        llist->tail->next=node;
    }
    llist->tail=node;
}
void insertatbig(LinkedList* llist, int node_data) {
    Node* NewNode = create(node_data);
    NewNode->next=llist->head;
    llist->head=NewNode;
}
void insertatend(LinkedList* llist, int node_data) {
    Node* NewNode = create(node_data);
    llist->tail->next=NewNode;
    llist->tail=NewNode;
}
void insertatany(LinkedList* llist, int node_data,int pos) {
    Node* NewNode = create(node_data);
    int i=0;
    Node* p=llist->head,*temp;
    while (i<pos-1)
    {
        p=p->next;
      i++;
    }
    NewNode->next=p->next;
    p->next=NewNode;
    
}
void DelectNodebig(LinkedList* llist){
    struct Node *Temp;
    Temp=llist->head;
    llist->head=Temp->next;
    free(Temp);
}
void DelectNodeend(LinkedList* llist){
    struct Node *t;
    t=llist->head;
    while (t->next->next!=NULL){
        t=t->next;
    }
    // printf("dec %d\n",t->data);
    llist->tail=t;
    t->next=NULL;
    free(t->next);
}
void DelectNodeany(LinkedList* llist,int pos){
    struct Node *t,*temp;
    t=llist->head;
    int i=0;
    while (i<pos-1){
        t=t->next;
        i++;
    }
    if (t->next->data==llist->tail->data){
        printf("out of range error  line number 95\n");
    }else{
        temp=t->next;
        t->next=temp->next;
    free(temp);
    }

}

int lenth(LinkedList* llist){
    int count=0;
    p=llist->head;
    while (p!=NULL){
         p=p->next;
         count++;
    }
    return count;
}

void reverse_list(LinkedList* llist){
    Node *curr,*next ,*prev=NULL,*temp;
    curr=llist->head;
    temp=llist->head;
    next=llist->head;
    while (next!=NULL){
        next=next->next;
        curr->next=prev;
        prev=curr;
        curr=next;
    } 
    llist->head=prev; 
    llist->tail=temp; 
}
void printLinkedList(LinkedList* llist) {
    struct Node *p;
    p=llist->head;
    printf("head: %d\n",llist->head->data);
    while (p!=NULL){
        printf("%d\n",p->data);
        p=p->next;
    }
    printf("tail: %d\n",llist->tail->data);
    
}

void reverse_list(LinkedList* llist){
    Node* temp=llist->head;
    llist->head=llist->tail;
    llist->tail=temp;
}


int main(){
    LinkedList* llist =(LinkedList*) malloc(sizeof(LinkedList));
    llist->head = NULL;
    llist->tail = NULL;
    // insert(&llist, 10);
    // insert(&llist, 20);
    // insert(&llist, 30);
    // insert(&llist, 40);
    // insert(&llist, 50);
    // insert(&llist, 60);
    insertat(llist,10);
    insertat(llist,20);
    insertat(llist,30);
    insertat(llist,40);
    insertat(llist,50);
    insertat(llist,60);
    insertat(llist,70);
    // insertatbig(llist,80);
    // insertatbig(llist,90);
    // insertatend(llist,100);
    // insertatend(llist,200);
    // insertatany(llist,300,3);
     // DelectNodeend(llist);
     // DelectNodeany(llist,);
     reverse_list(llist);
    printLinkedList(llist);
    reverse_list(llist);
    printLinkedList(llist);
//    int b= lenth(llist);
//    printf("%d\n",b);
    return 0;
}




























// #include <stdio.h>
// #include <stdlib.h>


// typedef struct node node;
// struct node
// {
//     int data;
//     node *next;   
// };
// node *head,*newNode,*tail ,*p,*temp;
// void  append(int data){
//     newNode= (node*)  malloc(sizeof(node));
//     newNode->data=data;
//     newNode->next = NULL;
//    if (head==NULL){
//     head= newNode ;
//    }else{
//     tail->next = newNode;
//    }
//    tail=newNode;
// }

// void display(){
//     p=head;
//     printf("%d\n",p->data);
//     while (p!=0)
//     {
//        printf("%d ",p->data);
//        p=p->next;
//     }
//     printf("\n");
//     printf("%d",tail->data);
// }
// void insertstart(int data){
//     newNode= (node*)  malloc(sizeof(node));
//     newNode->data=data;
//     newNode->next=head;
//     head=newNode;
// }
// void insertend(int data){
//      newNode= (node*) malloc(sizeof(node));
//      newNode->data=data;
//      newNode->next= NULL;
//     p=head;
//     while (p->next!=NULL){
//        p=p->next;
//     }
//     p->next=newNode;
//     tail=newNode;
// }
// void insertatpos(int data, int pos){
//     newNode= (node*) malloc(sizeof(node));
//     newNode->data=data;
//     int i=0;
//     p=head;
//     while(i<pos-1&&p->next!=NULL){
//         p=p->next;
//         i+=1;
//     }
//     newNode->next=p->next;
//     p->next=newNode;
// }


// void delectfrombig(){
//     temp=head;
//     head=temp->next;
//     free(temp);
// }
// void delectfromend(){
//    p=head;
//    while (p->next->next!=NULL)
//    {
//     p=p->next;
//    }
//    tail=p;
//    tail->next=NULL;
//    free(tail->next);
   
    
// }
// void delectatany(int pos){
//    p=head;int i=0;
//    while (i<pos-1){
//     p=p->next;
//     i++;
//    }
//    printf("%d",p->data);
//    temp=p->next;
//    p->next=temp->next;
//    free(temp);
    
// }
// int main(){
//     int a=5;
// append(10);
// append(20);
// append(30);
// append(40);
// insertstart(50);
// insertstart(30);
// insertstart(110);
// insertstart(40);
// insertend(80);
// insertatpos(60,8);
// delectfrombig();
// delectfromend();
// delectatany(3);
// display();
// printf("Hello, world!");
// printf("Hello, world!");
    
// }