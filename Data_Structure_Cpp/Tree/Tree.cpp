//
// Created by huang on 2024/6/13.
//

//��������ʵ��
#include "iostream"

using namespace std;

//������
template<class T>
struct TreeNode{
    T data;
    TreeNode * left;
    TreeNode * right;
    TreeNode(T x) : data(x),left(NULL),right(NULL){}

};

//������ģ����
template<class T>
class BinaryTree{
public:
    //���캯��
    BinaryTree() : root(NULL){}

    //��������
    ~BinaryTree();

    //���ø��ڵ�
    void Setroot(TreeNode<T>* &node){
        this->root = node;
    }

    //ǰ�����
    void preOrder(TreeNode<T>* node);
    void preOrder();

    //�������
    void inOrder(TreeNode<T>* node);
    void inOrder();

    //��������
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

// ��������
template<class T>
BinaryTree<T>::~BinaryTree() {
    destroyTree(root);
}

// ǰ�����
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

// �������
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

// �������
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

    cout << "ǰ��������: ";
    tree.preOrder();
    cout << endl;

    cout << "����������: ";
    tree.inOrder();
    cout << endl;

    cout << "�����������: ";
    tree.postOrder();
    cout << endl;

    delete root->left->left;
    delete root->left->right;
    delete root->left;
    delete root->right;
    delete root;

    return 0;
}