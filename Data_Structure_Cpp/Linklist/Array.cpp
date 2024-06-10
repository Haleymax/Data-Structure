//
// Created by huang on 2024/6/10.
//

//使用动态数组的思维实现

#include "iostream"
#define MaxSize 10

using namespace std;

typedef struct Arry{
    int *data;
    int length;  //线性表长度
}SeqList;


//初始化
void InitList(SeqList &L){
    L.data = (int *) malloc(MaxSize*sizeof (int ));
    L.length = 0;
}


//增加数组长度
void IncreaseSize(SeqList &L,int len){
    int *p = L.data;
    L.data = (int *) malloc((L.length + len)*sizeof (int ));   //新开辟一片空间
    for (int i = 0; i < L.length; ++i) {   //把旧数据移入新数组中
        L.data[i] = p[i];
    }
    free(p);
}

//按值查找
int LocateElem(SeqList L,int e){
    for (int i = 0; i < L.length; ++i) {
        if (L.data[i] == e)
            return i;   //查找成功返回索引
    }
    return -1;   //查找失败返回-1
}

//根据索引查找
int LocateIndex(SeqList L,int index){
    return L.data[index-1];
}

int main(){
    SeqList L;
    InitList(L);
    IncreaseSize(L,5);

    cout << LocateIndex(L,3) << endl;

    cout << LocateElem(L,2) << endl;


}