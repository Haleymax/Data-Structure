//
// Created by huang on 2024/6/13.
//

/**
 * �����ڽӾ���ʵ�ֵ�ͼ������ͼ��
 * **/
#ifndef DATA_STRUCTURE_CPP_AM_GRAPH_H
#define DATA_STRUCTURE_CPP_AM_GRAPH_H

#endif //DATA_STRUCTURE_CPP_AM_GRAPH_H

#pragma once

#include "iostream"
#define MaxSize 20
using namespace std;

class Graph{
private:
    int v;  //������
    int arc[MaxSize][MaxSize];
    bool read[MaxSize];
public:

    //���캯��
    Graph(int num);

    //��ӱ�
    bool addEdge(int start , int end);

    //��ӡͼ
    void printGraph();

    //�����������
    void BFS(int start);

    //�����������
    void DFS();
};


Graph::Graph(int num) {
    if (num > MaxSize) {
        cout << "������ͼ����󶥵���" << endl;
        exit(1);
    }

    //��ʼ���ڽӾ���
    for (int i = 0; i < num; ++i) {
        for (int j = 0; j < num; ++j) {
            this->arc[i][j] = 0;   //���ʾû�б�
        }
        this->read[i] = false;
    }

    this->v = num;
}

//��ӱߣ�����ͼ��
bool Graph::addEdge(int start , int end) {
    if (start > MaxSize || end > MaxSize){
        cout << "������Χ" << endl;
        return false;
    }
    this->arc[start][end] = 1;
    this->arc[end][start] = 1;
    return true;
}

//��ӡͼ
void Graph::printGraph(){
    for (int i = 0; i < this->v; ++i) {
        for (int j = 0; j < this->v; ++j) {
            cout << arc[i][j] << " ";
        }
        cout << endl;
    }
}