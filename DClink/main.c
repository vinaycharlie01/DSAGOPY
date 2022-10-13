#include <stdio.h>
#include <stdlib.h>

typedef struct Node Node;
typedef struct Link Link;
struct Node{
    int data;
    Node* next;
    Node* prev;
};
struct Link{
    Node* head;
    Node* tail;
};

Node* createNode(int data){
    Node* node = (Node*)malloc(sizeof(Node));
    node->data = data;
    node->next = NULL;
    node->prev = NULL;
    return node;
}

void append(Link* llist,int data ){
    Node* NewNode = createNode(data);
    if(llist->head==NULL){
        llist->head = NewNode;
        llist->head->next=NewNode;
    }else{
        llist->tail->next = NewNode;
        NewNode->prev=llist->tail;
        NewNode->next=llist->head;
    }
    llist->tail = NewNode;
    llist->head->prev=NewNode;
}

void displayreverce(Link* llist){
    Node* temp = llist->tail;
    while(temp!=llist->head){
        printf("%d ",temp->data);
        temp = temp->prev;
    }
    printf("%d ",temp->data);
}
void displayforword(Link* llist){
    Node* temp = llist->head;
    while(temp!=llist->tail){
        printf("%d ",temp->data);
        temp = temp->next;
    }
    printf("%d ",temp->data);
}

void insertats(Link* llist,int data){
    Node* NewNode = createNode(data);
    NewNode->next=llist->head;
    llist->head->prev=NewNode;
    NewNode->prev=llist->tail;
    llist->head=NewNode;
    llist->tail->next=NewNode;
}
void insertatany(Link* llist,int data,int pos){
    Node* NewNode = createNode(data);
    Node* p=llist->head;int i=0;
    while(i<pos-1){
        p=p->next;
        i++;
    }
    NewNode->next=p->next;
    NewNode->prev=p;
    p->next->prev=NewNode;
    p->next=NewNode;
    //printf("%d",p->data);
}


void delectfrombig(Link* llist){
    Node* temp = llist->head;
    llist->head=temp->next;
    temp->next->prev=llist->tail;
    free(temp);
}
void delectfromend(Link* llist){
    Node* temp = llist->tail;
    llist->tail=temp->prev;
    temp->prev->next=llist->head;
    llist->head->prev=temp;
    free(temp);

}
void delectfromany(Link* llist ,int pos){
    Node* temp = llist->head;
    int i=0;
    while(i<pos-1){
        temp=temp->next;
        i++;
    }
    // temp->prev->next=temp->next;
    // temp->next->prev=temp->prev;
    // free(temp);
    Node* p=temp->next;
    temp->next=temp->next->next;
    temp->next->next->prev=temp;
    free(p);

}
Link* New(Link* l){
    l = (Link*)malloc(sizeof(Link));
    l->head = NULL;
    l->tail = NULL;
    return l;
}

int main(){
    // Link* llist=(Link*) malloc(sizeof(Link));
    // llist->head=NULL;
    // llist->tail=NULL;
    Link* llist;
    llist = New(NULL);
    append(llist,1);
    append(llist,2);
    append(llist,3);
    append(llist,4);
    append(llist,5);
    append(llist,6);
   // insertats(llist,7);
   // delectfrombig(llist);
  //  delectfromend(llist);
    delectfromany(llist, 2);
  //insertatany(llist,8,4);
    //reverce_list(llist);
    displayforword(llist);
    printf("\n");
    // displayreverce(llist);
    return 0;
}