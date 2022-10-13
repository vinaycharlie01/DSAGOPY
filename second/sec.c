// #include <assert.h>
// #include <limits.h>
// #include <math.h>
// #include <stdbool.h>
// #include <stddef.h>
// #include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
// #include <string.h>

char* readline();

typedef struct SinglyLinkedListNode SinglyLinkedListNode;
typedef struct SinglyLinkedList SinglyLinkedList;

struct SinglyLinkedListNode {
    int data;
    SinglyLinkedListNode* next;
};

struct SinglyLinkedList {
    SinglyLinkedListNode* head;
    SinglyLinkedListNode* tail;
};

SinglyLinkedListNode* create_singly_linked_list_node(int node_data) {
    SinglyLinkedListNode* node = (SinglyLinkedListNode*)malloc(sizeof(SinglyLinkedListNode));
    node->data = node_data;
    node->next = NULL;
    return node;
}

<<<<<<< HEAD
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
=======
void insert_node_into_singly_linked_list(int node_data) {
    SinglyLinkedListNode* node = create_singly_linked_list_node(node_data);
    SinglyLinkedList* llist = (SinglyLinkedList*)malloc(sizeof(SinglyLinkedList));
    llist->head = NULL;
    llist->tail = NULL;
    if (llist->head==NULL) {
        llist->head = node;
    } else {
        llist->tail->next = node;
>>>>>>> c86f2e85f11047d04ecda404bed9389244dd9c77
    }

    llist->tail = node;
}

void free_singly_linked_list(SinglyLinkedListNode* node) {
    while (node) {
        SinglyLinkedListNode* temp = node;
        node = node->next;

        free(temp);
    }
}
void printLinkedList() {
    SinglyLinkedList* llist = (SinglyLinkedList*)malloc(sizeof(SinglyLinkedList));
    SinglyLinkedListNode* p=llist->head;
    printf("head: %d\n",p->data);
    while (p!=NULL){
        printf("%d\n",(p->data));
        p=p->next;
    }
   printf("tail: %d\n",p->data);
}


int main(){
    insert_node_into_singly_linked_list(10);
    insert_node_into_singly_linked_list(20);
    insert_node_into_singly_linked_list(20);
    insert_node_into_singly_linked_list(20);
    //printLinkedList();
    return 0;
}

