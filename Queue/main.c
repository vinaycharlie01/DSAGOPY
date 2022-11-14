#include <stdio.h>

int N =5;
int queue [5];
int front =-1;
int rear =-1;
void enqueue(int x){
    if (rear == N-1){
        printf("Overflow");
    }else if (front == -1&& rear == -1){
        front=rear=0,
        queue[rear]=x;
    }else{
        rear++;
        queue[rear]=x;
    }
}


void delect(){
    if (front == -1 && rear == -1){
        printf("list is empty"); 
    }else if (front==rear){
        front=rear=-1;
    }else{
        front++;
    }
}

void display(){
    if (front==-1 && rear==-1){
        printf("list is empty");
    }else
    {
        int p =front;
        while (p!=rear+1)
        {
            printf("%d  ",queue[p]);
            p++;
        }
    }   
}


void top(){
    if (front==-1 && rear==-1){
        printf("list is empty");
    }else{
        printf("%d ",queue[front]);
    }
}

int main(){
    enqueue(10);
    enqueue(20);
    enqueue(40);
    enqueue(50);
    delect();
    display();
    printf("\n");
    top();
    return 0;
}
