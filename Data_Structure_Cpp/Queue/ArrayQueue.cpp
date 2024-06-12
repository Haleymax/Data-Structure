//
// Created by huang on 2024/6/12.
//

#include "iostream"
#define MaxSize 10
using namespace std;

//采用模板以便处理多种数据类型(这里T的类型主要为简单的类型，不做复杂处理)
template<class T>
class Queue{

public:
    //构造函数初始化队列
    Queue() : front(0) , rear(0) ,len(0){}

    //入队操作
    bool Enqueue(T v);

    //出队操作
    bool Dequuq(T &e);

    //检查队列是否为空
    bool IsEmpty();

    //获取队列大小
    int Size();

    //检查队列是否已满
    bool IsFull();

    //打印队列
    void Show();

private:
    T data[MaxSize];  //队列本体
    int front;  //队头
    int rear;   //队尾
    int len;  //队长
};

//入队
template<class T>
bool Queue<T>::Enqueue(T v) {
    if (IsFull()){
        return false;
    }

    this->data[rear] = v;  //从队尾入队
    this->rear = (this->rear + 1) %MaxSize;
    this->len++;
    return true;
}

//出队
template<class T>
bool Queue<T>::Dequuq(T &e) {
    if (IsEmpty()){
        return false;
    }

    e = this->data[front];  //从对头出队
    this->front = (this->front + 1) % MaxSize;
    this->len--;
    return true;
}

//检查队列是否为空
template<class T>
bool Queue<T>::IsEmpty() {
    if (this->len == 0)
        return true;
    return false;
}

//检查队列是否已满
template<class T>
bool Queue<T>::IsFull() {
    if (this->len == MaxSize)
        return true;
    return false;
}

//获取队列大小
template<class T>
int Queue<T>::Size() {
    return this->len;
}

//打印队列
template<class T>
void Queue<T>::Show() {
    if (IsEmpty()) {
        cout << " Is empty" << endl;
        return;
    }
    cout<<"--------------"<<endl;
    int index = front;
    while (index != rear){
        cout<< data[index] << " ";
        index = (index + 1)%MaxSize;
    }
    cout << endl;
    cout<<"--------------"<<endl;

}


int main(){
    Queue<int> q;
    Queue<char> c;
    int e;
    char x;


    //q的入队
    q.Enqueue(1);
    q.Enqueue(2);
    q.Enqueue(3);
    q.Enqueue(4);
    q.Show();
    q.Dequuq(e);
    cout << e << endl;
    q.Show();

    c.Enqueue('a');
    c.Enqueue('b');
    c.Enqueue('c');
    c.Enqueue('d');
    c.Show();
    c.Dequuq(x);
    cout << x << endl;
    c.Show();
}