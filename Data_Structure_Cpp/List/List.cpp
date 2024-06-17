//
// Created by huang on 2024/6/17.
//

//采用双向列表来实现一个类似于STL中的list容器

#include "iostream"
using namespace std;


//双向链表节点
template<class T>
struct ListNode{
    T data;
    ListNode* pre;
    ListNode* next;

    ListNode(T x) : data(x), pre(NULL) , next(NULL){}
};

//双向链表模板类的实现
template<class T>
class List{
private:
    ListNode<T> *head;
    ListNode<T> *tail;

    size_t size;

public:

    //构造实现初始化
    List() : head(NULL),tail(NULL),size(0){}

    //析构清空
    ~List(){
        clear();
    }

    //清空链表
    void clear(){
        while (head){
            ListNode<T> *temp = head;
            head = head->next;
            delete temp;
        }
        head = tail = NULL;
        size = 0;
    }

    //插入元素到链表尾部
    void push_back(T x){
        ListNode<T> *newNode = new ListNode<T>(x);
        if (!head)
            head = tail = newNode;   //如果链表为空则这就是第一个节点
        else{
            tail->next = newNode;
            newNode->pre = tail;
            tail = newNode;
        }
        size++;
    }

    //插入元素到链表头部
    void push_front(T x){
        ListNode<T>* newNode = new ListNode<T>(x);
        if (!head){
            head = tail = newNode;
        } else{
            newNode->next = head;
            head->pre = newNode;
            head = newNode;
        }
        size++;
    }

    //插入元素到指定位置
    void insert(size_t pos , T x){
        if (pos > size){
            cout<<" 超出列表范围 " << endl;
            return;
        }
        ListNode<T>* newNode = new ListNode<T>(x);
        if (size == 0){
            head = tail = newNode;
        } else if (pos == 0){
            push_front(x);
        } else if (pos == size){
            push_back(x);
        } else{
            ListNode<T>* current = head;
            for (size_t i = 0; i < pos-1; ++i) {
                current = current->next;
            }
            newNode->next = current->next;
            current->next->pre = newNode;
            current->next = newNode;
            newNode->pre = current;
            size++;
        }
    }

    //重载运算符 =
    List &operator=(const List& other){
        if (this != &other){
            clear();
            for (ListNode<T>*current = other.head; current ; current = current->next) {
                push_back(current->data);
            }
        }
        return *this;
    }

    //重载一个list容器中没有的运算
    T& operator[](size_t pos) {
        if (pos >= size) {
            throw std::out_of_range("Index is out of range");
        }
        ListNode<T>* current = head;
        for (size_t i = 0; i < pos; ++i) {
            current = current->next;
        }
        return current->data;
    }

    //显示元素
    void show(){
        ListNode<T> *current = head;
        while (current){
            cout << current->data << " ";
            current = current->next;
        }
        cout << endl;
    }
};

int main(){
    List<int> list;
    list.push_back(1);
    list.push_back(2);
    list.push_back(3);
    list.push_back(4);
    list.push_back(5);
    list.show();

    list.insert(3,9);
    list.show();


    cout << list[3]<<endl;

    list.clear();
    list.show();

    return 0;
}

