//
// Created by huang on 2024/6/26.
//

#include "iostream"
#include "vector"

//�����ϣ���С
const int TABLE_SIZE = 128;

using namespace std;

//��ϣ������
class MyHashSet{
private:

    //��ϣ��
    vector<int> table;

    //��¼Ԫ������
    int size;

    //��Ԫ������
    int hashFunction(int key){
        return key % TABLE_SIZE;
    }

public:
    //���캯������ʼ����ϣ��
    MyHashSet() : table(TABLE_SIZE,-1) , size(0){}

    //����Ԫ��
    void insert(int key);

    //ɾ��Ԫ��
    void remove(int key);

    //����Ԫ���Ƿ����
    bool contains(int key);

    //��ȡ��ǰ���ϵĴ�С
    int getSize() const;
};

//����Ԫ��
void MyHashSet::insert(int key) {
    int hash = hashFunction(key);   //��ȡλ������
    int originalHash = hash;   //��¼�������λ��

    //����ͻ����������̽�⴦���ͻ
    while (table[hash] != -1 && table[hash] != key){
        hash = (hash + 1) % TABLE_SIZE;  //������ͻ�������ƶ�һλ

        //���ò�����չ�Ĺ�ϣ��������˾Ͳ��ܹ�������
        if (hash == originalHash){
            cout << "��ϣ���������޷�������Ԫ��" << endl;
            return;
        }
    }

    //����ҵ���λ�þͽ�Ԫ�ز���
    if (table[hash] == -1){
        table[hash] = key;
        size++;
    }
}

//ɾ��Ԫ��
void MyHashSet::remove(int key) {
    int hash = hashFunction(key);
    int orignalHash = hash;

    //����̽�����Ԫ��
    while (table[hash] != -1){
        if (table[hash] == key){
            //�ҵ�Ԫ�أ�ɾ��
            table[hash] = -1;
            size--;
            return;
        }
        hash = (hash + 1) % TABLE_SIZE;

        //���̽��ص�ԭ����λ�ô����������Լ������Լ�û�����Ԫ��
        if (hash == orignalHash)
            return;
    }
}

//����Ԫ���Ƿ����
bool MyHashSet::contains(int key) {
    int hash = hashFunction(key);
    int originalHash = hash;

    //����̽�ⷨ
    while (table[hash] != -1){
        if (table[hash] == key)
            return true;

        hash = (hash + 1) % TABLE_SIZE;

        //���̽���ԭ��λ�ñ�ʾ������
        if (hash == originalHash)
            return false;
    }

    return false;
}

int MyHashSet::getSize() const {
    return size;
}

// �����Զ����ϣ����
int main() {
    MyHashSet mySet;
    mySet.insert(1);
    mySet.insert(2);
    mySet.insert(3);

    std::cout << "�������Ƿ����2: " << (mySet.contains(2) ? "��" : "��") << std::endl;
    std::cout << "���ϴ�С: " << mySet.getSize() << std::endl;

    mySet.remove(2);
    std::cout << "ɾ��2�󣬼������Ƿ����2: " << (mySet.contains(2) ? "��" : "��") << std::endl;
    std::cout << "���ϴ�С: " << mySet.getSize() << std::endl;

    return 0;
}