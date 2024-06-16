//
// Created by huang on 2024/6/13.
//

//基于邻接表实现的图的广度优先搜索和深度优先搜索

#include "AL_graph.hpp"
#include "queue"
#include "iostream"
using namespace std;

//广度优先搜索
void Graph::BFS(int start) {
    queue<int> q;
    q.push(start);
    read[start] = true;

    while (!q.empty()){
        int v = q.front();
        q.pop();
        cout << v << " ";

        for (vector<int>::iterator it = adj[v].begin();it != adj[start].end() ; it++) {
            if (!read[*it]){
                q.push(*it);
                read[*it] = true;
            }
        }
    }
}

//DFS （时间有限，采用递归实现了）
void Graph::DFS(int start) {
    read[start] = true;
    cout << start << " ";
    for (vector<int>::iterator it = adj[start].begin() ; it != adj[start].end() ; it++) {
        if (!read[*it]){
            DFS(*it);
        }
    }

}

int main() {
    Graph g(5);

    g.addEdge(0, 1);
    g.addEdge(0, 4);
    g.addEdge(1, 2);
    g.addEdge(1, 3);
    g.addEdge(1, 4);
    g.addEdge(2, 3);
    g.addEdge(3, 4);

    cout << "DFS: ";
    g.DFS(0);
    cout << endl;

    cout << "BFS: ";
    g.BFS(0);
    cout << endl;

    return 0;
}