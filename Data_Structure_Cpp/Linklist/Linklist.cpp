//
// Created by huang on 2024/6/10.
//

#include "iostream"

using namespace std;

typedef struct LinkNode{
    int Data;
    LinkNode *Next;
}LinkNode,*Linklist;

int main(){

    //节点1
    Linklist node = new LinkNode{2, NULL};

    //节点2
    Linklist node1 = new LinkNode{3, NULL};
    node->Next = node1;

    //节点三
    Linklist node2 = new LinkNode{5, NULL};
    node1->Next = node2;

    //打印节点
    Linklist l = node;
    while (l != NULL){
        cout << l->Data << endl;
        l = l->Next;
    }

    //释放节点
    while (node != NULL){
        Linklist temp = node;
        node = node->Next;
        delete node;
    }
}
