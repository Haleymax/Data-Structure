//
// Created by huang on 2024/6/13.
//
/**
 * 基于邻接表实现的图（无向图）
 * **/

#ifndef DATA_STRUCTURE_CPP_AL_GRAPH_H
#define DATA_STRUCTURE_CPP_AL_GRAPH_H

#endif //DATA_STRUCTURE_CPP_AL_GRAPH_H

#pragma once

#include "iostream"
#include "vector"
#include "stack"
#define MaxSize 10

using namespace std;

class Graph{
public:
    int v;
    vector<vector<int>> adj; //使用vector容器来实现邻接表了（想自己写链表来实现的但又感觉太麻烦了）
    bool read[MaxSize];

public:

    //构造函数
    Graph(int num);

    //添加边
    bool addEdge(int start , int end);

    //光度优先搜索
    void BFS(int start);

    //深度优先搜索
    void DFS(int start);
};

//构造函数初始化图
Graph::Graph(int num) {
    if (num < 0 || num > MaxSize){
        cout << "超出范围 "<< endl;
        exit(1);
    }
    this->v = num;

    adj.resize(num);  //初始化邻接表
    for (int i = 0; i < num; ++i) {
        this->read[i] = false;
    }
}

//添加边
bool Graph::addEdge(int start, int end) {
    if (start > MaxSize || end > MaxSize){
        cout << "超出范围 "<< endl;
        return false;
    }
    adj[start].push_back(end);
    if (start != end){
        adj[end].push_back(start);
    }
    return true;
}

