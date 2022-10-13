#include <stdio.h>
#include <stdlib.h>

typedef struct Node Node;
typedef struct Clink Clink;
struct Node{
    int data;
    Node* next;
};

struct Clink{
    Node* head;
};


Node* CreateNode(int data){
    Node* node = (Node*)malloc(sizeof(Node));
    node->data = data;
    node->next = NULL;
    return node;
}

void append(Clink* llist,int data){
    Node* node = CreateNode(data);
    Node* temp;
    if(llist->head == NULL){
        llist->head = node;
    }else{
        temp->next=node;
    }
    temp=node;
    node->next=llist->head;
}


void print(Clink* llist){
    printf("[ ");
    Node* temp= llist->head;
    while (temp->next!=llist->head){
        printf("%d ",temp->data);
        temp=temp->next;
    }
    printf("%d",temp->data);
    printf(" ]\n");
    //printf("%d",temp->next->data);
}

void pushbig(Clink* llist,int data){
    Node* node = CreateNode(data);
    Node* temp=llist->head;
    while(temp->next!=llist->head){
        temp=temp->next;
    }
    node->next=llist->head;
    llist->head=node;
    temp->next=llist->head;
}
void pushend(Clink* llist,int data){
    Node* node = CreateNode(data);
    Node* temp=llist->head;
    while(temp->next!=llist->head){
        temp=temp->next;
    }
    temp->next=node;
    node->next=llist->head;
}

int len(Clink* llist){
    int count=1;
    Node* p=llist->head;
    while(p->next!=llist->head){
       p=p->next;
       count++;
    }
    if (llist->head==NULL){
        return 0;
    }
    return count;
}
void pushany(Clink* llist,int data,int pos){
    Node* node = CreateNode(data);
    int i=0;
    Node* p=llist->head;
    while(i<pos-1){
        p=p->next;
        i++;
    }
    Node* temp=p->next;
    p->next=node;
    node->next=temp;
}


void  pop(Clink* llist,int pos){
    if (llist->head==NULL){
        free(llist->head);
        printf("Your list is enpty...!");
    }
   if (pos==0){
     Node* p=llist->head;
    Node* temp= llist->head;
    while (temp->next!=p){
        temp=temp->next;
    }
   llist->head=p->next;
   temp->next=llist->head;
//    printf("%d\n", llist->head->data);
//    printf("%d\n",temp->next->data);
   }else if (pos>0&&pos<len(llist)){
       int i=0;
       Node* tem= llist->head;
       while(i<pos-1){
        tem=tem->next;
           i++;
       }
       Node* temp=tem->next;
       tem->next=tem->next->next;
       free(temp);
   }else{
      Node* temp= llist->head;
      while (temp->next->next!=llist->head){
        temp=temp->next;
      }
      ///printf("%d\n",temp->next->data);
      Node* q=temp->next;
      temp->next=llist->head;
      free(q);
   }
   }

void reverse(Clink* llist){
    Node* temp=llist->head;
    while(temp->next!=llist->head){
        temp=temp->next;
    }
    Node* prev=temp;
    Node* curr=llist->head;
    Node* next=llist->head;
    while(next->next!=llist->head){
        next=next->next;
        curr->next=prev;
        prev=curr;
        curr=next;
    }
    curr->next=prev;
    llist->head=curr;
}

Clink* New(Clink* l){
    l=(Clink*)malloc(sizeof(Clink));
    l->head=NULL;
    return l;
}

int main(){
    Clink* llist = New(NULL);
    append(llist,1);
    append(llist,2);
    append(llist,3);
    append(llist,4);
    append(llist,5);
    append(llist,6);
    append(llist,7);
    append(llist,8);
    append(llist,9);
    append(llist,10);
    append(llist,11);
    append(llist,12);
    // pushbig(llist,10);
    // pushbig(llist,20);
    // pushend(llist,40);
   // pushany(llist,30,8);
    //printf("%d\n", len(llist));
    pop(llist,2);
    reverse(llist);
     pop(llist,2);

   print(llist);
    printf("\n");
}