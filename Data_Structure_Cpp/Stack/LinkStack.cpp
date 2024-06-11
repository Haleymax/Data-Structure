//
// Created by huang on 2024/6/11.
//

#include "iostream"

using namespace std;

typedef struct Node{
    int data;
    Node *Next;
}Node;

class Stack{
private:
    Node *top;

public:

    //构造函数
    Stack();

    //析构函数
    ~Stack();

    //判空
    bool IsEmpty();

    //进栈
    bool Push(int e);

    //出栈
    bool Pop(int &e);

    //获取栈顶元素
    bool GetTop(int &e);

    //打印栈
    void Show();

};

//构造函数初始化栈
Stack::Stack() {
    this->top = NULL;
}

//析构函数用于销毁栈
Stack::~Stack() {
    int e;
    while (!IsEmpty()){
        this->Pop(e);
    }
}

//判空
bool Stack::IsEmpty() {
    if (top == NULL)
        return true;
    return false;
}

//进栈
bool Stack::Push(int e){
    Node *newNode = new Node {e,top};
    if (newNode == NULL)
        return false;
    top = newNode;
    return true;
}

//出栈
bool Stack::Pop(int &e) {
    if (IsEmpty())
        return false;
    Node *temp = top;  //创建工作指针保留栈顶节点防止断链
    e = top->data;
    top = top->Next;
    delete temp;
    return true;
}

//获取栈顶元素
bool Stack::GetTop(int &e) {
    if (IsEmpty())
        return false;
    e = top->data;
    return true;
}

void Stack::Show() {
    Node *r = top;
    while (r){
        cout << r->data << endl;
        r = r->Next;
    }
    cout << "--------------" << endl;
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
