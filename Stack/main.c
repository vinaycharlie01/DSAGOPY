// #include <stdio.h>
// #include <stdlib.h>
// #define N 7

// int stack[N];


// int top=-1;
// void push(int data){
//     if(top==N-1){
//         printf("Stack Overflow\n");
//     }else{
//         ++top;
//         stack[top]=data;
//     }
// }

// void pop(){
//     int item;
//     if(top ==-1){
//         printf("Stack Underflow\n");
//     }else{
//         item = stack[top];
//         stack[top]=0;
//         top--;

//     }
// }
// void Top(){
//     if(top == -1){
//         printf("Stack is empty\n");
//     }else{
//         printf("Top element is %d\n",stack[top]);
//     }
// }

// void isempty(){
//     if (top == -1){
//         printf("Stack is empty\n");
//     }
// }

// void display(){
//     for (int i=top;i>=0;i--) {
//         printf("%d ",stack[i]);
//     }
// }

// int main(){
//    push(10);
//    push(20);
//    push(30);
//    push(40);
//    push(50);
//    push(60);
//    push(70);
//    display();
//    pop();
//    display();
//    printf("\n");
//    printf("%d\n",stack[6]);
//    return 0;
// }

#include <stdio.h>
#include <stdlib.h>


typedef struct Node Node;
typedef struct link Link;
struct Node{
    int data;
    Node* next;
};

struct link{
    Node* head;
};


Node* createNode(int data){
    Node* node = (Node*)malloc(sizeof(Node));
    node->data = data;
    node->next = NULL;
    return node;
}


void  Push(Link* link, int data){
    Node* newNode = createNode(data);
    if (link->head==NULL){
        link->head = newNode;
    }else{
        newNode->next = link->head;
        link->head = newNode;
    }
}

void  Pop(Link* link){
    Node* temp = link->head;
    link->head = link->head->next;
    free(temp);
}

void  Bottom(Link* link){
    Node* temp = link->head;
    while(temp->next!=NULL){
        temp = temp->next;
    }
    printf("%d\n",temp->data);
}

void  Top(Link* link){
    Node* temp = link->head;
    printf("%d\n",temp->data);
}

void isempty(Link* link){
    if (link->head == NULL){
        printf("Stack is empty\n");
    }else{
        printf("Top element is %d\n",link->head->data);
    }
}
void Display(Link* link){
    Node* temp = link->head;
    while(temp!=NULL){
        printf("%d ",temp->data);
        temp = temp->next;
    }
}

int main(){
    Link* llist=(Link*) malloc(sizeof(Link));
    llist->head=NULL;
    Push(llist,10);
    Push(llist,20);
    Push(llist,30);
    Push(llist,40);
    Push(llist,50);
    Push(llist,60);
    Push(llist,70);
    Push(llist,80);
    Pop(llist);
    Pop(llist);
    Display(llist);
    printf("\n");
   
}
