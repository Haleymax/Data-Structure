//
// Created by huang on 2024/6/13.
//

#include "AM_graph.hpp"
#include "queue"    //懒得写队列和栈了，先用下STL中现成的
#include "stack"


//广度优先搜索
void Graph::BFS(int strat) {
    queue<int> q;
    this->read[strat] = true;
    q.push(strat);

    cout << strat<< " ";
    while (!q.empty()){
        int v = q.front();  //每次取队头元素若是有没有被访问过的邻接节点就将其入队
        q.pop();

        for (int i = 0; i < this->v; ++i) {
            if (arc[strat][i] == 1 && !read[i]){
                q.push(i);
                read[i] = true;
                cout << i << " ";
            }
        }
    }
}

//深度优先搜索
void Graph::DFS(int start) {
    stack<int> s;
    s.push(start);
    this->read[start] = true;
    while (!s.empty()){
        int v = s.top();
        cout << v << " ";
        s.pop();

        for (int i = 0; i < MaxSize; ++i) {
            int neighbor = arc[v][i];
            if (!read[neighbor]){
                s.push(neighbor);
                read[neighbor] = true;
            }
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

    cout << "BFS : ";
    g.BFS(0);
    cout << "DFS : ";
    g.DFS(0);
    return 0;
}