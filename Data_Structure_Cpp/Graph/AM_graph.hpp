//
// Created by huang on 2024/6/13.
//

/**
 * 基于邻接矩阵实现的图（无向图）
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
    int v;  //顶点数
    int arc[MaxSize][MaxSize];
    bool read[MaxSize];
public:

    //构造函数
    Graph(int num);

    //添加边
    bool addEdge(int start , int end);

    //打印图
    void printGraph();

    //广度优先搜索
    void BFS(int start);

    //深度优先搜索
    void DFS();
};


Graph::Graph(int num) {
    if (num > MaxSize) {
        cout << "超出了图的最大顶点数" << endl;
        exit(1);
    }

    //初始化邻接矩阵
    for (int i = 0; i < num; ++i) {
        for (int j = 0; j < num; ++j) {
            this->arc[i][j] = 0;   //零表示没有边
        }
        this->read[i] = false;
    }

    this->v = num;
}

//添加边（无向图）
bool Graph::addEdge(int start , int end) {
    if (start > MaxSize || end > MaxSize){
        cout << "超出范围" << endl;
        return false;
    }
    this->arc[start][end] = 1;
    this->arc[end][start] = 1;
    return true;
}

//打印图
void Graph::printGraph(){
    for (int i = 0; i < this->v; ++i) {
        for (int j = 0; j < this->v; ++j) {
            cout << arc[i][j] << " ";
        }
        cout << endl;
    }
}