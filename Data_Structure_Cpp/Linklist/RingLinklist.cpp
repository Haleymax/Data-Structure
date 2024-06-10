//
// Created by huang on 2024/6/10.
//

//循环单链表的实现
#include "iostream"
using namespace std;

typedef struct LNode{
    int data;
    LNode *Next;
}LNode,*Linklist;

//初始化
bool InitList(Linklist &L){
    L = (LNode *) malloc(sizeof (LNode));
    if (L == NULL)
        return false;
    L->Next = L;
    return true;
}

//添加新元素
Linklist Apped(Linklist l , int e){
    LNode *s = (LNode *) malloc(sizeof (LNode));
    s->data = e;
    LNode *tail;
    tail = l;
    while (tail->Next != l)
        tail = tail->Next;
    tail->Next = s;
    s->Next = l;
    return l;
}


//删除某一元素
bool Del(Linklist l,int e){
    if (l == NULL)
        return false;
    LNode *r = l->Next;
    LNode *pre = l;
    while (r != l){
        if (r->data == e){
            pre->Next = r->Next;
            free(r);
            return true;
        }
        pre = r;
        r = r->Next;
    }
    return false;
}

// 显示循环链表元素
void DisplayList(Linklist L) {
    if (L == NULL) {
        cout << "The list is empty." << endl;
        return;
    }

    LNode *current = L->Next; // 从第一个数据节点开始
    if (current == L) {
        cout << "The list contains only the header node." << endl;
        return;
    }

    do {
        cout << current->data << " "; // 打印当前节点的数据
        current = current->Next; // 移动到下一个节点
    } while (current != L); // 当回到头节点时停止

    cout << endl; // 打印换行符以美化输出
}

int main(){
    Linklist l;
    InitList(l);
    l = Apped(l, 1);
    DisplayList(l);
    l = Apped(l, 2);
    l = Apped(l, 3);
    Del(l,2);
    DisplayList(l);


    return 0;
}