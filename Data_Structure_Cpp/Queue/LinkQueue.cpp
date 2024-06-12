//
// Created by huang on 2024/6/12.
//

#include "iostream"

using namespace std;

//定义链表节点
template<class T>
struct Node{
    T data;
    Node* Next;
};

//链表队列
template<class T>
class Queue{
public:
    //构造函数
    Queue():front(NULL),rear(NULL){}

    //入队
    void Enqueue(T v);

    //出队
    bool Dequeue(T &e);

    //检查队列是否为空
    bool IsEmpty();

    //获取队列大小
    int Size();

    //打印队列
    void Show();

    //析构函数释放空间
    ~Queue();

private:
    Node<T> *front;   //队头指针
    Node<T> *rear;    //队尾指针
    int len;          //兑长
};

//入队
template<class T>
void Queue<T>::Enqueue(T v){
    Node<T> *newNode = new Node<T>;
    newNode->data = v;
    newNode->Next = NULL;
    if (rear){
        rear->Next = newNode;  //对尾入队
        rear = newNode;
    } else
        front = rear = newNode;  //链表为空
    this->len++;
}

//出队
template<class T>
bool Queue<T>::Dequeue(T &e){
    if (IsEmpty())
        return false;
    Node<T> *temp = front;  //保存头节点防止锻链
    e = front->data;
    front = front->Next;
    this->len--;
    return true;
}

//检查队列是否为空
template<class T>
bool Queue<T>::IsEmpty(){
    if (this->front == NULL && this->front == this->rear)
        return true;
    return false;
}

//获取队列大小
template<class T>
int Queue<T>::Size(){
    return this->len;
}

//打印队列
template<class T>
void Queue<T>::Show(){
    if (IsEmpty()) {
        cout << "Queue is empty" << endl;
        return;
    }
    cout << "--------------" << endl;
    Node<T>* temp = front;
    while (temp) {
        cout << temp->data << " ";
        temp = temp->Next;
    }
    cout << endl << "--------------" << endl;
}

//析构函数释放空间
template<class T>
Queue<T>::~Queue(){
    // 遍历链表并逐个删除节点
    Node<T>* current = front;
    while (current != nullptr) {
        Node<T>* temp = current;
        current = current->Next;
        delete temp;
    }
    front = rear = nullptr; // 将头尾指针置空
}

int main() {
    Queue<int> q;
    int e;

    // q的入队
    q.Enqueue(1);
    q.Enqueue(2);
    q.Enqueue(3);
    q.Enqueue(4);
    q.Show();
    q.Dequeue(e);
    cout << e << endl;
    q.Show();

    return 0;
}