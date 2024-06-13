//
// Created by huang on 2024/6/13.
//

//二叉树的层次遍历
#include <iostream>
using namespace std;

// 二叉树节点
template<class T>
struct TreeNode {
    T data;
    TreeNode *left;
    TreeNode *right;
    TreeNode(T x) : data(x), left(nullptr), right(nullptr) {}
};

// 队列节点
template<class T>
struct QueueNode {
    TreeNode<T>* treeNode;
    QueueNode<T>* next;
    QueueNode(TreeNode<T>* node) : treeNode(node), next(nullptr) {}
};

// 队列结构体
template<class T>
class Queue {

public:
    QueueNode<T>* front; // 队首
    QueueNode<T>* rear;  // 队尾

    Queue() : front(nullptr), rear(nullptr) {}

    // 入队
    void enqueue(TreeNode<T>* node) {
        QueueNode<T>* newNode = new QueueNode<T>(node);
        if (rear) {
            rear->next = newNode;
            rear = newNode;
        } else {
            front = rear = newNode;
        }
    }

    // 出队
    TreeNode<T>* dequeue() {
        if (!front) return nullptr;
        QueueNode<T>* temp = front;
        TreeNode<T>* data = front->treeNode;
        front = front->next;
        if (front == nullptr) rear = nullptr;
        delete temp;
        return data;
    }

    // 判断队列是否为空
    bool isEmpty() {
        return front == nullptr;
    }
};

// 二叉树模板类
template<class T>
class BinaryTree {
public:
    // 构造函数
    BinaryTree() : root(nullptr) {}

    //设置根节点
    void Setroot(TreeNode<T>* &node){
        this->root = node;
    }

    void levelOrder() {
        Queue<T> q;
        if (root != nullptr) q.enqueue(root);

        while (!q.isEmpty()) {
            TreeNode<T>* node = q.dequeue();
            cout << node->data << " ";
            if (node->left != nullptr) q.enqueue(node->left);
            if (node->right != nullptr) q.enqueue(node->right);
        }
    }

private:
    TreeNode<T> *root;
};

int main() {
    BinaryTree<int> tree;

    TreeNode<int>* root = new TreeNode<int>(1);
    root->left = new TreeNode<int>(2);
    root->right = new TreeNode<int>(3);
    root->left->left = new TreeNode<int>(4);
    root->left->right = new TreeNode<int>(5);

    tree.Setroot(root);

    cout << "二叉树的层次遍历结果: ";
    tree.levelOrder();
    cout << endl;


    return 0;
}
