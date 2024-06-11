//
// Created by huang on 2024/6/11.
//

#include "iostream"
#include "stdlib.h"
#define MaxSize 10

using namespace std;

class Stack{
private:
    int *data;
    int top;  //栈顶索引

public:

    //构造函数
    Stack();

    //析构函数
    ~Stack();

    //判空
    bool IsEmpty();

    //求表长
    int GetLength();

    //进栈
    bool Push(int x);

    //出栈
    bool Pop(int &e);

    //读取栈顶元素
    bool GetTop(int &e);

    //打印栈中的元素
    void Show();

};

//构造函数初始化栈
Stack::Stack(){
    data = new int[MaxSize];  //构造函数构建栈的数组
    this->top = -1;
}

//析构函数用于销毁栈
Stack::~Stack() {
    delete[] data;
}

//判空
bool Stack::IsEmpty() {
    if (this->top == -1)
        return true;
    return false;
}

//求栈的元素个数
int Stack::GetLength() {

    return this->top + 1;
}

//进栈
bool Stack::Push(int x) {
    if (this->top == MaxSize)
        return false;    //表示栈中的元素已满了不能够再进栈了
    this->top++;
    this->data[this->top] = x;
    return true;
}

//出栈
bool Stack::Pop(int &e) {
    if (this->top == -1)
        return false;   //表示栈中已经没有元素了
    e = this->data[this->top];
    this->top--;
    return true;
}

//读取栈顶元素
bool Stack::GetTop(int &e) {
    if (this->top == -1)
        return false;
    e = this->data[this->top];
    return true;
}

//打印栈中的元素
void Stack::Show() {
    if (this->top <= 0){
        cout<< "The stack is empty" << endl;
        return;
    }
    for (int i = 0; i < this->top+1; ++i) {
        cout << this->data[i] << endl;
    }
    cout << "----------------------------" << endl;
}


int main(){
    Stack s;
    int e;
    s.Push(1);
    s.Push(2);
    s.Push(3);
    s.Show();
    s.GetTop(e);
    cout<<  e  <<endl;
    s.Pop(e);
    cout<<  e  <<endl;
}

