//
// Created by huang on 2024/6/13.
//

//二叉树的实现
#include "iostream"

using namespace std;

//二叉树
template<class T>
struct TreeNode{
    T data;
    TreeNode * left;
    TreeNode * right;
    TreeNode(T x) : data(x),left(NULL),right(NULL){}

};

//二叉树模板类
template<class T>
class BinaryTree{
public:
    //构造函数
    BinaryTree() : root(NULL){}

    //析构函数
    ~BinaryTree();

    //设置根节点
    void Setroot(TreeNode<T>* &node){
        this->root = node;
    }

    //前序遍历
    void preOrder(TreeNode<T>* node);
    void preOrder();

    //中序遍历
    void inOrder(TreeNode<T>* node);
    void inOrder();

    //后续遍历
    void postOrder(TreeNode<T>* node);
    void postOrder();


private:
    TreeNode<T> *root;

    void destroyTree(TreeNode<T>* &node){
        if (node == NULL)
            return;
        destroyTree(node->left);
        destroyTree(node->right);
        delete node;
        node = NULL;
    }
};

// 析构函数
template<class T>
BinaryTree<T>::~BinaryTree() {
    destroyTree(root);
}

// 前序遍历
template<class T>
void BinaryTree<T>::preOrder(TreeNode<T>* node) {
    if (node == NULL) return;
    cout << node->data << " ";
    preOrder(node->left);
    preOrder(node->right);
}

template<class T>
void BinaryTree<T>::preOrder() {
    preOrder(root);
}

// 中序遍历
template<class T>
void BinaryTree<T>::inOrder(TreeNode<T>* node) {
    if (node == NULL) return;
    inOrder(node->left);
    cout << node->data << " ";
    inOrder(node->right);
}

template<class T>
void BinaryTree<T>::inOrder() {
    inOrder(root);
}

// 后序遍历
template<class T>
void BinaryTree<T>::postOrder(TreeNode<T>* node) {
    if (node == NULL) return;
    postOrder(node->left);
    postOrder(node->right);
    cout << node->data << " ";
}

template<class T>
void BinaryTree<T>::postOrder() {
    postOrder(root);
}

int main() {
    BinaryTree<int> tree;

    TreeNode<int>* root = new TreeNode<int>(1);
    root->left = new TreeNode<int>(2);
    root->right = new TreeNode<int>(3);
    root->left->left = new TreeNode<int>(4);
    root->left->right = new TreeNode<int>(5);

    tree.Setroot(root);

    cout << "前序遍历结果: ";
    tree.preOrder();
    cout << endl;

    cout << "中序遍历结果: ";
    tree.inOrder();
    cout << endl;

    cout << "后续遍历结果: ";
    tree.postOrder();
    cout << endl;

    delete root->left->left;
    delete root->left->right;
    delete root->left;
    delete root->right;
    delete root;

    return 0;
}