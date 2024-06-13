//
// Created by huang on 2024/6/13.
//

//�������Ĳ�α���
#include <iostream>
using namespace std;

// �������ڵ�
template<class T>
struct TreeNode {
    T data;
    TreeNode *left;
    TreeNode *right;
    TreeNode(T x) : data(x), left(nullptr), right(nullptr) {}
};

// ���нڵ�
template<class T>
struct QueueNode {
    TreeNode<T>* treeNode;
    QueueNode<T>* next;
    QueueNode(TreeNode<T>* node) : treeNode(node), next(nullptr) {}
};

// ���нṹ��
template<class T>
class Queue {

public:
    QueueNode<T>* front; // ����
    QueueNode<T>* rear;  // ��β

    Queue() : front(nullptr), rear(nullptr) {}

    // ���
    void enqueue(TreeNode<T>* node) {
        QueueNode<T>* newNode = new QueueNode<T>(node);
        if (rear) {
            rear->next = newNode;
            rear = newNode;
        } else {
            front = rear = newNode;
        }
    }

    // ����
    TreeNode<T>* dequeue() {
        if (!front) return nullptr;
        QueueNode<T>* temp = front;
        TreeNode<T>* data = front->treeNode;
        front = front->next;
        if (front == nullptr) rear = nullptr;
        delete temp;
        return data;
    }

    // �ж϶����Ƿ�Ϊ��
    bool isEmpty() {
        return front == nullptr;
    }
};

// ������ģ����
template<class T>
class BinaryTree {
public:
    // ���캯��
    BinaryTree() : root(nullptr) {}

    //���ø��ڵ�
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

    cout << "�������Ĳ�α������: ";
    tree.levelOrder();
    cout << endl;


    return 0;
}
